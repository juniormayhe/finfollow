package main

import (
	"fmt"
	"net/http"
	"strconv"
)

// Define a home handler function which writes a byte slice containing
// "Hello from FinFollow" as the response body.
func home(w http.ResponseWriter, r *http.Request) {
	// Check if the current request URL path exactly matches "/". If it doesn't
	// the http.NotFound() function to send a 404 response to the client.

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		// Importantly, we then return from the handler. If we don't return the hand
		// would keep executing and also write the "Hello from SnippetBox" message.
		return
	}

	w.Write([]byte("Hello from FinFollow"))

}

// Add a showAsset handler function.
func showAsset(w http.ResponseWriter, r *http.Request) {
	// Extract the value of the id parameter from the query string and try to
	// convert it to an integer using the strconv.Atoi() function. If it can't
	// be converted to an integer, or the value is less than 1, we return a 404
	// not found response.
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	//w.Write([]byte("SHOW asset..."))
	fmt.Fprintf(w, "SHOW asset ID... %d", id)
}

// Add a addAsset handler function.
// test it: curl -i -X POST http://localhost:4000/asset/add -d "name=test&description=test"
func addAsset(w http.ResponseWriter, r *http.Request) {
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
		http.Error(w, "Method Not Allowed", 405)
		return
	}

	w.Write([]byte("ADD asset..."))
}
