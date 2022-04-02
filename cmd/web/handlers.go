package main

import (
	"fmt"
	"net/http"
	"time"

	"juniormayhe.com/finfollow/pkg/forms"
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

	// Create an instance of a templateData struct holding the slice of assets
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

	// Use the PopString() method to retrieve the value for the "flash" key.
	// PopString() also deletes the key and value from the session data, so it
	// acts like a one-time fetch. If there is no matching key in the session
	// data this will return the empty string.

	// moved to helpers.go
	// flash := app.session.PopString(r, "flash")

	data := &templateData{
		Asset: asset,
		//Flash: flash, moved to helpers.go
	}

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
	// data (a models.Asset struct) as the final parameter.
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

// Add a new addAssetForm handler, which for now returns a placeholder res
func (app *application) addAssetForm(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("Create a new asset..."))
	app.renderTemplate(w, r, "create.page.gohtml", &templateData{
		Now: time.Now().Format("2006-01-02"),
		// Pass a new empty forms.Form object to the template.
		Form: forms.New(nil),
	})
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

	// name := "Polygon"
	// value := float64(8.2)
	// currency := "MATIC"
	// custody := "Metamask"
	// created := time.Now()
	//
	//

	// First we call r.ParseForm() which adds any data in POST request bodies
	// to the r.PostForm map. This also works in the same way for PUT and PATCH
	// requests. If there are any errors, we use our app.ClientError helper to
	// a 400 Bad Request response to the user.
	err := r.ParseForm()
	if err != nil {
		app.infoLog.Println("Error parsing form")
		app.clientError(w, http.StatusBadRequest)
		return
	}

	// Create a new forms.Form struct containing the POSTed data from the
	// form, then use the validation methods to check the content.
	form := forms.New(r.PostForm)
	form.Required("name", "value", "currency", "custody", "created")
	form.MaxLength("name", 100)
	form.ValidNumber("value")
	form.MaxLength("currency", 10)
	form.MaxLength("custody", 50)
	// form.PermittedValues("expires", "365", "7", "1")

	// If the form isn't valid, redisplay the template passing in the
	// form.Form object as the data.
	if !form.Valid() {
		app.renderTemplate(w, r, "create.page.gohtml", &templateData{
			Now:  time.Now().Format("2006-01-02"),
			Form: form})
		return
	}

	// Use the r.PostForm.Get() method to retrieve the relevant data fields
	// from the r.PostForm map.
	// name := strings.TrimSpace(r.PostForm.Get("name"))
	// value, errValue := strconv.ParseFloat(strings.TrimSpace(r.PostForm.Get("value")), 64)
	// currency := strings.TrimSpace(r.PostForm.Get("currency"))
	// custody := strings.TrimSpace(r.PostForm.Get("custody"))
	// created := time.Now()

	finished, _ := time.Parse("YYYY-MM-DD", "0001-01-01")
	active := true

	// Initialize a map to hold any validation errors.
	//errors := make(map[string]string)
	// if name == "" {
	// 	errors["name"] = "Please enter a name"
	// } else if utf8.RuneCountInString(name) > 100 {
	// 	errors["name"] = "The name is too long (maximum is 100 characters)"
	// }

	// if r.PostForm.Get("value") == "" {
	// 	errors["value"] = "Value is required"
	// } else if value < 0 {
	// 	errors["value"] = "Please enter a value greater or equal to 0"
	// } else if errValue != nil {
	// 	errors["value"] = "Value is invalid"
	// }

	// if currency == "" {
	// 	errors["currency"] = "Please enter a currency"
	// } else if utf8.RuneCountInString(name) > 10 {
	// 	errors["currency"] = "The currency is too long (maximum is 10 characters)"
	// }

	// if custody == "" {
	// 	errors["custody"] = "Please enter a custody"
	// } else if utf8.RuneCountInString(name) > 50 {
	// 	errors["custody"] = "The custody is too long (maximum is 50 characters)"
	// }

	// var errCreated error = nil
	// if r.PostForm.Get("created") == "" {
	// 	errors["created"] = "Please enter a date"
	// } else {

	// 	// 2006-01-02 layout has constants standing for long year 2006, zero month 01, zero day 02
	// 	created, errCreated = time.ParseInLocation("2006-01-02", strings.TrimSpace(r.PostForm.Get("created")), time.UTC)

	// 	if errCreated != nil {
	// 		errors["custody"] = "Date format is invalid"
	// 	}
	// }

	// If there are any errors, dump them in a plain text HTTP response and ret
	// from the handler.
	// if len(errors) > 0 {
	// 	// fmt.Fprint(w, errors)
	// 	app.renderTemplate(w, r, "create.page.gohtml", &templateData{
	// 		FormErrors: errors,
	// 		FormData:   r.PostForm,
	// 	})

	// 	return
	// }

	id, dbErr := app.model.Insert("wander", form.Get("name"), form.GetNumber("value"), form.Get("currency"), form.Get("custody"), form.GetDate("created"), finished, active)

	//w.Write([]byte("SHOW asset..."))
	if dbErr != nil {
		app.infoLog.Printf("Error while inserting in Database: %s", dbErr)
		app.serverError(w, dbErr)
		return
	}

	app.infoLog.Printf("Added asset with id = %s", id)
	_, balanceErr := app.model.UpdateBalance("wander", form.GetNumber("value"))
	if balanceErr != nil {
		app.serverError(w, balanceErr)
		return
	}

	// w.Write([]byte("ADD asset..."))

	// Use the Put() method to add a string value ("Your asset was saved
	// successfully!") and the corresponding key ("flash") to the session
	// data. Note that if there's no existing session for the current user
	// (or their session has expired) then a new, empty, session for them
	// will automatically be created by the session middleware.
	app.session.Put(r, "flash", "Asset successfully added!")

	// Redirect the user to the relevant page for the asset.
	// http.Redirect(w, r, fmt.Sprintf("/asset?id=%v", id), http.StatusSeeOther)

	// Change the redirect to use the new semantic URL style of /asset/:id
	http.Redirect(w, r, fmt.Sprintf("/asset/%s", id), http.StatusSeeOther)
}
