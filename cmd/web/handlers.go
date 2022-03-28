package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

// Define a home handler function to render home page
// we inject dependencies into handler passing the application struct value
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	// Check if the current request URL path exactly matches "/". If it doesn't
	// the http.NotFound() function to send a 404 response to the client.

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		// Importantly, we then return from the handler. If we don't return the hand
		// would keep executing and also write the "Hello from SnippetBox" message.
		return
	}

	// Initialize a slice containing the paths to the two files. Note that the
	// home.page.tmpl file must be the *first* file in the slice.
	files := []string{
		"./ui/html/home.page.gohtml",
		"./ui/html/base.layout.gohtml",
		"./ui/html/footer.partial.gohtml",
	}

	//w.Write([]byte("Hello from FinFollow"))

	// Use the template.ParseFiles() function to read the template file into a
	// template set. Notice that we can pass the slice of file
	// as a variadic parameter
	ts, err := template.ParseFiles(files...)
	if err != nil {
		// If there's an error, we log the detailed error message and
		// the http.Error() function to send a generic 500 Internal Server Error
		// response to the user.
		// log.Println(err.Error())
		// Because the home handler function is now a method against application
		// it can access its fields, including the error logger. We'll write the
		// message to this instead of the standard logger.
		// app.errorLog.Println(err.Error())
		// http.Error(w, "Internal Server Error", 500)

		app.serverError(w, err) // this is our helper method in helpers.go
		return
	}

	// We then use the Execute() method on the template set to write the template
	// content as the response body. The last parameter to Execute() represents
	// dynamic data that we want to pass in, which for now we'll leave as nil.
	err = ts.Execute(w, nil)
	if err != nil {
		// log.Println(err.Error())
		// Also update the code here to use the error logger from the applicatio
		// struct.
		// app.errorLog.Println(err.Error())
		// http.Error(w, "Internal Server Error", 500)
		app.serverError(w, err)
	}

}

// Add a showAsset handler function.
func (app *application) showAsset(w http.ResponseWriter, r *http.Request) {
	// Extract the value of the id parameter from the query string and try to
	// convert it to an integer using the strconv.Atoi() function. If it can't
	// be converted to an integer, or the value is less than 1, we return a 404
	// not found response.
	id := r.URL.Query().Get("id")
	if len(id) <= 0 {
		// http.NotFound(w, r)
		app.notFound(w) // use the notFound helper method in helpers.go
		return
	}

	asset, err := app.assets.Get(id)
	if err != nil {
		app.serverError(w, err)
		return
	}

	fmt.Fprintf(w, "\nasset = %+v", asset)

}

// Add a addAsset handler function.
// test it: curl -i -X POST http://localhost:4000/asset/add -d "name=test&description=test"
func (app *application) addAsset(w http.ResponseWriter, r *http.Request) {
	// Use r.Method to check whether the request is using POST or not.
	// If it's not, use the w.WriteHeader() method to send a 405 status code and
	// the w.Write() method to write a "Method Not Allowed" response body. We
	// then return from the function so that the subsequent code is not executed
	if r.Method != "POST" {

		// If user is using wrong method, let user know POST is the only method allowed:
		// Use the Header().Set() method to add an 'Allow: POST' header to the
		// response header map. The first parameter is the header name, and
		// the second parameter is the header value.
		w.Header().Set("Allow", "POST")

		// w.WriteHeader(405)
		// w.Write([]byte("Method Not Allowed"))
		// http.Error(w, "Method Not Allowed", 405)
		app.clientError(w, http.StatusMethodNotAllowed) // use the clientError helper method in helpers.go
		return
	}

	name := "Cardano"
	value := float32(50.0)
	currency := "ADA"
	custody := "Binance"
	created := time.Now()
	finished, _ := time.Parse("YYYY-MM-DD", "0000-00-00")
	active := true
	id, dbErr := app.assets.Insert(name, value, currency, custody, created, finished, active)

	//log.Println(ref.Path)
	//w.Write([]byte("SHOW asset..."))
	if dbErr != nil {

		app.serverError(w, dbErr)
		return
	}

	// Redirect the user to the relevant page for the asset.
	http.Redirect(w, r, fmt.Sprintf("/asset?id=%s", id), http.StatusSeeOther)

	//w.Write([]byte("ADD asset..."))
}
