package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

// The serverError helper writes an error message and stack trace to the errorLogger
// then sends a generic 500 Internal Server Error response to the user.
func (app *application) serverError(w http.ResponseWriter, err error) {
	// append error message and append stack trace of the current goroutine with debug.Stack()
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())

	// set the frame depth to one step back in stack trace
	// to get the filename and line number where the error happened
	// ie: if handlers.go is the caller, it would be printed handlers.go instead of helpers.go
	app.errorLog.Output(2, trace)

	app.errorLog.Println(trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// The clientError helper sends a specific status code and corresponding description
// to the user. We'll use this later in the book to send responses like 400 "Bad
// Request" when there's a problem with the request that the user sent.
func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

// For consistency, we'll also implement a notFound helper. This is simply a
// convenience wrapper around clientError which sends a 404 Not Found response
// the user.
func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}
