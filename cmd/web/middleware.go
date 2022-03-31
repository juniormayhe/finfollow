// middleware functions here must be used in routes.go
package main

import (
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
