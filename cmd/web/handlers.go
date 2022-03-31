package main

import (
	"fmt"
	"net/http"
	"time"

	"juniormayhe.com/finfollow/pkg/models"
)

// Define a home handler function to render home page
// we inject dependencies into handler passing the application struct value
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	// Check if the current request URL path exactly matches "/". If it doesn't
	// the http.NotFound() function to send a 404 response to the client.
	// Because Pat matches the "/" path exactly, we can now remove the manual c
	// of r.URL.Path != "/" from this handler.
	// if r.URL.Path != "/" {
	// 	app.notFound(w)
	// 	// Importantly, we then return from the handler. If we don't return the hand
	// 	// would keep executing and also write the "Hello from SnippetBox" message.
	// 	return
	// }

	assets, err := app.model.Latest("wander")
	if err != nil {
		app.serverError(w, err)
		return
	}

	// for _, asset := range assets {
	// 	fmt.Fprintf(w, "%+v\n", asset)
	// }

	// Create an instance of a templateData struct holding the slice of
	// snippets.
	data := &templateData{Assets: assets}

	// Initialize a slice containing the paths to the two files. Note that the
	// home.page.tmpl file must be the *first* file in the slice.
	// files := []string{
	// 	"./ui/html/home.page.gohtml",
	// 	"./ui/html/base.layout.gohtml",
	// 	"./ui/html/footer.partial.gohtml",
	// }

	//w.Write([]byte("Hello from FinFollow"))

	// Use the template.ParseFiles() function to read the template file into a
	// template set. Notice that we can pass the slice of file
	// as a variadic parameter
	//ts, err := template.ParseFiles(files...)
	//if err != nil {
	// If there's an error, we log the detailed error message and
	// the http.Error() function to send a generic 500 Internal Server Error
	// response to the user.
	// log.Println(err.Error())
	// Because the home handler function is now a method against application
	// it can access its fields, including the error logger. We'll write the
	// message to this instead of the standard logger.
	// app.errorLog.Println(err.Error())
	// http.Error(w, "Internal Server Error", 500)

	//	app.serverError(w, err) // this is our helper method in helpers.go
	//	return
	//}

	// We then use the Execute() method on the template set to write the template
	// content as the response body. The last parameter to Execute() represents
	// dynamic data that we want to pass in, which for now we'll leave as nil.
	//err = ts.Execute(w, data)
	// if err != nil {
	// log.Println(err.Error())
	// Also update the code here to use the error logger from the applicatio
	// struct.
	// app.errorLog.Println(err.Error())
	// http.Error(w, "Internal Server Error", 500)

	//	app.serverError(w, err)
	//}

	app.renderTemplate(w, r, "home.page.gohtml", data)

}

// Add a showAsset handler function.
func (app *application) showAsset(w http.ResponseWriter, r *http.Request) {
	// Extract the value of the id parameter from the query string and try to
	// convert it to an integer using the strconv.Atoi() function. If it can't
	// be converted to an integer, or the value is less than 1, we return a 404
	// not found response.
	// id := r.URL.Query().Get("id")

	// Pat doesn't strip the colon from the named capture key, so we need to
	// get the value of the named capture ":id" from the
	// query string instead of "id".
	id := r.URL.Query().Get(":id")
	if len(id) <= 0 {
		// http.NotFound(w, r)
		app.notFound(w) // use the notFound helper method in helpers.go
		return
	}

	asset, err := app.model.Get("wander", id)

	if err != nil {
		if err == models.ErrNoRecord {
			app.notFound(w)
			return
		}
		app.serverError(w, err)
		return
	}

	data := &templateData{Asset: asset}

	// Initialize a slice containing the paths to the show.page.tmpl file,
	// plus the base layout and footer partial that we made earlier.
	// files := []string{
	// 	"./ui/html/show.page.gohtml",

	// 	"./ui/html/base.layout.gohtml",
	// 	"./ui/html/footer.partial.gohtml",
	// }

	// Parse the template files...
	// ts, err := template.ParseFiles(files...)
	// if err != nil {
	// 	app.serverError(w, err)
	// 	return
	// }
	// And then execute them. Notice how we are passing in the snippet
	// data (a models.Snippet struct) as the final parameter.
	// err = ts.Execute(w, asset)
	// Pass in the templateData struct when executing the template.
	// err = ts.Execute(w, data)
	// if err != nil {
	// 	app.serverError(w, err)
	// }

	// Use the new render helper.
	app.renderTemplate(w, r, "show.page.gohtml", data)

	app.infoLog.Printf("asset = %+v", asset)

}

// Add a new createSnippetForm handler, which for now returns a placeholder res
func (app *application) addAssetForm(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a new asset..."))
}

// Add a addAsset handler function.
// test it: curl -i -X POST http://localhost:4000/asset/add -d "name=test&description=test"
func (app *application) addAsset(w http.ResponseWriter, r *http.Request) {
	// Use r.Method to check whether the request is using POST or not.
	// If it's not, use the w.WriteHeader() method to send a 405 status code and
	// the w.Write() method to write a "Method Not Allowed" response body. We
	// then return from the function so that the subsequent code is not executed
	// With Pat the check of r.Method != "POST" is now superfluous
	// and can be removed.
	// if r.Method != "POST" {
	// 	// If user is using wrong method, let user know POST is the only method allowed:
	// 	// Use the Header().Set() method to add an 'Allow: POST' header to the
	// 	// response header map. The first parameter is the header name, and
	// 	// the second parameter is the header value.
	// 	w.Header().Set("Allow", "POST")

	// 	// w.WriteHeader(405)
	// 	// w.Write([]byte("Method Not Allowed"))
	// 	// http.Error(w, "Method Not Allowed", 405)
	// 	app.clientError(w, http.StatusMethodNotAllowed) // use the clientError helper method in helpers.go
	// 	return
	// }

	name := "Polygon"
	value := float32(8.2)
	currency := "MATIC"
	custody := "Metamask"
	created := time.Now()
	finished, _ := time.Parse("YYYY-MM-DD", "1980-01-01")
	active := true
	id, dbErr := app.model.Insert("wander", name, value, currency, custody, created, finished, active)

	//w.Write([]byte("SHOW asset..."))
	if dbErr != nil {
		app.serverError(w, dbErr)
		return
	}
	app.infoLog.Printf("Added asset with id = %s", id)
	_, balanceErr := app.model.UpdateBalance("wander", value)
	if balanceErr != nil {
		app.serverError(w, balanceErr)
		return
	}

	// w.Write([]byte("ADD asset..."))

	// Redirect the user to the relevant page for the asset.
	// http.Redirect(w, r, fmt.Sprintf("/asset?id=%v", id), http.StatusSeeOther)

	// Change the redirect to use the new semantic URL style of /snippet/:id
	http.Redirect(w, r, fmt.Sprintf("/asset/%s", id), http.StatusSeeOther)
}
