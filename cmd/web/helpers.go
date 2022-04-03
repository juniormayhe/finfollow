package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/justinas/nosurf"
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

// Create an addDefaultData helper. This takes a pointer to a templateData
// struct, adds the current year to the CurrentYear field, and then returns
// the pointer. Again, we're not using the *http.Request parameter at the
// moment, but we will do later in the book.
func (app *application) addDefaultData(td *templateData, r *http.Request) *templateData {

	if td == nil {
		td = &templateData{}
	}

	td.AuthenticatedUser = app.authenticatedUser(r)
	log.Println(">>>> AuthenticatedUser:", td.AuthenticatedUser)
	td.CurrentYear = time.Now().Year()

	// Add the flash message to the template data, if one exists.
	// message is automatically included the next time any page is rendered
	td.Flash = app.session.PopString(r, "flash")

	// Add the CSRF token to the templateData struct.
	td.CSRFToken = nosurf.Token(r)

	return td
}

func (app *application) renderTemplate(w http.ResponseWriter, r *http.Request, name string, td *templateData) {
	// Retrieve the appropriate template set from the cache based on the page n
	// (like 'home.page.tmpl'). If no entry exists in the cache with the
	// provided name, call the serverError helper method that we made earlier.
	ts, ok := app.templateCache[name]
	if !ok {
		app.serverError(w, fmt.Errorf("The template %s does not exist", name))
		return
	}
	// Execute the template set, passing in any dynamic data.
	// if the template cannot be executed go will return 200 with encoded html content
	// err := ts.Execute(w, td)
	// if err != nil {
	// 	app.serverError(w, err)
	// }

	// dynamic data at app level to be rendered in template
	defaultData := app.addDefaultData(td, r)

	// Initialize a new buffer.
	buf := new(bytes.Buffer)

	// try to write the template to the buffer, instead of straight to the
	// http.ResponseWriter. If there's an error, call our serverError helper
	// and return.
	err := ts.Execute(buf, defaultData)
	if err != nil {
		app.serverError(w, err)
		return
	}

	// Write the contents of the buffer to the http.ResponseWriter. Again, this
	// is another time where we pass our http.ResponseWriter to a function that
	// takes an io.Writer.
	buf.WriteTo(w)

}

// The authenticatedUser method returns the ID of the current user from the
// session, or zero if the request is from an unauthenticated user.
func (app *application) authenticatedUser(r *http.Request) string {
	return app.session.GetString(r, "userID")
}
