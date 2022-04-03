// middleware functions here must be used in routes.go
package main

import (
	"fmt"
	"net/http"
)

// middleware to act on every request that is received,
// we need it to be executed before a request hits our servemux
// secureHeaders → servemux → application handler
func secureHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Any code here will execute on the way down the chain.
		// secureHeaders → servemux → application handler
		w.Header().Set("X-XSS-Protection", "1; mode=block")
		w.Header().Set("X-Frame-Options", "deny")

		// if user is not authorized you can send 403 forbidden status code
		// and early return before ServeHTTP to stop executing the chain

		next.ServeHTTP(w, r)
		// Any code here will execute on the way back up the chain.
		// application handler → servemux → secureHeaders

	})
}

//logRequest ↔ secureHeaders ↔ servemux ↔ application handler
func (app *application) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.infoLog.Printf("%s - %s %s %s", r.RemoteAddr, r.Proto, r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}

// when a goroutine panics (an single http request), server will output an empty response.
// this middleware will recover the panic and call app.serverError to output 500 Internal Server Error
func (app *application) recoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Create a deferred function (which will always be run in the event
		// of a panic as Go unwinds the stack).
		defer func() {
			// Use the builtin recover function to check if there has been a
			// panic or not. If there has...
			if err := recover(); err != nil {
				// Set a "Connection: close" header on the response
				w.Header().Set("Connection", "close")

				// normalize the interface{} parameter passed to panic (string, error, etc) to an error
				normalizedError := fmt.Errorf("%s", err)

				// Call the app.serverError helper method to return a 500
				// Internal Server response.
				app.serverError(w, normalizedError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func (app *application) requireAuthenticatedUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// If the user is not authenticated, redirect them to the login page and
		// return from the middleware chain so that no subsequent handlers in
		// the chain are executed.
		if app.authenticatedUser(r) == "" {
			http.Redirect(w, r, "/user/login", http.StatusFound)
			return
		}
		// Otherwise call the next handler in the chain.
		next.ServeHTTP(w, r)
	})
}
