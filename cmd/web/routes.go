package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/justinas/alice" // package to create composable middleware
)

//func (app *application) routes() *http.ServeMux {
// use http.Handler instead of *http.ServeMux.
func (app *application) routes() http.Handler {

	// Create a middleware chain with justinas/alice
	// containing our 'standard' middleware
	// which will be used for every request our application receives.
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	// Create a new middleware chain containing the middleware specific to
	// our dynamic application routes. For now, this chain will only contain
	// the session middleware but we'll add more to it later.
	dynamicMiddleware := alice.New(app.session.Enable)

	// Use the http.NewServeMux() function to initialize a new servemux, then
	// register the home function as the handler for the "/" URL pattern.
	// mux := http.NewServeMux()
	// mux.HandleFunc("/", app.home)
	// mux.HandleFunc("/asset", app.showAsset)
	// mux.HandleFunc("/asset/add", app.addAsset)
	// since native Go doesnt support method based routing (GET, POST,..)
	// and doesn't support clean URLs, we need to use a custom router
	mux := pat.New()
	// mux.Get("/", http.HandlerFunc(app.home))
	// mux.Get("/asset/add", http.HandlerFunc(app.addAssetForm))
	// mux.Post("/asset/add", http.HandlerFunc(app.addAsset))
	// mux.Get("/asset/:id", http.HandlerFunc(app.showAsset)) // Moved down to give preference to exact match route before wildcard route
	// Update these routes to use the new dynamic middleware chain followed
	// by the appropriate handler function.
	mux.Get("/", dynamicMiddleware.ThenFunc(app.home)) // without alice: mux.Get("/", app.session.Enable(http.HandlerFunc(app.home)))
	// Add the requireAuthenticatedUser middleware to the chain.
	mux.Get("/asset/add", dynamicMiddleware.Append(app.requireAuthenticatedUser).ThenFunc(app.addAssetForm)) // without alice: mux.Get("/", app.session.Enable(app.requireAuthenticatedUser(http.HandlerFunc(app.home))))
	// Add the requireAuthenticatedUser middleware to the chain
	mux.Post("/asset/add", dynamicMiddleware.Append(app.requireAuthenticatedUser).ThenFunc(app.addAsset))
	mux.Get("/asset/:id", dynamicMiddleware.ThenFunc(app.showAsset))

	// Add the five new routes.
	mux.Get("/user/signup", dynamicMiddleware.ThenFunc(app.signupUserForm))
	mux.Post("/user/signup", dynamicMiddleware.ThenFunc(app.signupUser))

	mux.Get("/user/login", dynamicMiddleware.ThenFunc(app.loginUserForm))
	mux.Post("/user/login", dynamicMiddleware.ThenFunc(app.loginUser))

	// Add the requireAuthenticatedUser middleware to the chain.
	mux.Post("/user/logout", dynamicMiddleware.Append(app.requireAuthenticatedUser).ThenFunc(app.logoutUser))

	// Create a file server which serves files out of the "./ui/static" directo
	// Note that the path given to the http.Dir function is relative to the pro
	// directory root.
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// Use the mux.Handle() function to register the file server as the handler
	// all URL paths that start with "/static/". For matching paths, we strip the
	// "/static" prefix before the request reaches the file server.
	//mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))

	//return mux

	// Pass the servemux as the 'next' parameter to the secureHeaders middleware
	// Because secureHeaders is just a function, and the function returns a
	// http.Handler we don't need to do anything else.
	// return http.Handler instead of *http.ServeMux.
	// return secureHeaders(mux)

	// Wrap the existing chain with the logRequest middleware.
	// return app.logRequest(secureHeaders(mux))

	// Wrap the existing chain with the recoverPanic middleware.
	// return app.recoverPanic(app.logRequest(secureHeaders(mux)))

	// Return the 'standard' middleware chain
	// followed by the servemux.
	return standardMiddleware.Then(mux)
}
