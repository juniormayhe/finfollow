package main

import "net/http"

//func (app *application) routes() *http.ServeMux {
// use http.Handler instead of *http.ServeMux.
func (app *application) routes() http.Handler {

	// Use the http.NewServeMux() function to initialize a new servemux, then
	// register the home function as the handler for the "/" URL pattern.
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/asset", app.showAsset)
	mux.HandleFunc("/asset/add", app.addAsset)

	// Create a file server which serves files out of the "./ui/static" directo
	// Note that the path given to the http.Dir function is relative to the pro
	// directory root.
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// Use the mux.Handle() function to register the file server as the handler
	// all URL paths that start with "/static/". For matching paths, we strip the
	// "/static" prefix before the request reaches the file server.
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	//return mux

	// Pass the servemux as the 'next' parameter to the secureHeaders middleware
	// Because secureHeaders is just a function, and the function returns a
	// http.Handler we don't need to do anything else.
	// return http.Handler instead of *http.ServeMux.
	// return secureHeaders(mux)

	// Wrap the existing chain with the logRequest middleware.
	return app.logRequest(secureHeaders(mux))
}
