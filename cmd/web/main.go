package main

import (
	"context"
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/firestore"
	"juniormayhe.com/finfollow/pkg/firestoredb"
)

// Define an application struct to hold the application-wide dependencies for t
// web application. For now we'll only include fields for the two custom logger
// we'll add more to it as the build progresses.
type application struct {
	errorLog      *log.Logger
	infoLog       *log.Logger
	model         *firestoredb.FirestoreModel
	templateCache map[string]*template.Template
}

func main() {

	// Define a new command-line flag with the name 'addr', a default value of
	// and some short help text explaining what the flag controls. The value of
	// flag will be stored in the addr variable at runtime.
	// go run cmd/web/* -addr=":9999"
	addr := flag.String("addr", ":4000", "HTTP network address")

	// Importantly, we use the flag.Parse() function to parse the command-line
	// This reads in the command-line flag value and assigns it to the addr
	// variable. You need to call this *before* you use the addr variable
	// otherwise it will always contain the default value of ":4000". If any error
	// encountered during parsing the application will be terminated.
	flag.Parse()

	// Use log.New() to create a logger for writing information messages. This
	// three parameters: the destination to write the logs to (os.Stdout), a st
	// prefix for message (INFO followed by a tab), and flags to indicate what
	// additional information to include (local date and time). Note that the fl
	// are joined using the bitwise OR operator |.
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	// Create a logger for writing error messages in the same way, but use stderr
	// the destination and use the log.Lshortfile flag to include the relevant
	// file name and line number.
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.LUTC|log.Llongfile)

	client, err := openClient()
	if err != nil {
		errorLog.Fatalf("error initializing app: %v\n", err)
	}

	// Initialize a new template cache...
	templateCache, err := newTemplateCache("./ui/html/")
	if err != nil {
		errorLog.Fatal(err)
	}

	// Initialize a new instance of application containing the dependencies.
	app := &application{
		errorLog:      errorLog,
		infoLog:       infoLog,
		model:         &firestoredb.FirestoreModel{Client: client},
		templateCache: templateCache,
	}

	// We also defer a call to client.Close(), so that the connection is closed
	// before the main() function exits
	defer client.Close()

	// Initialize a new http.Server struct. We set the Addr and Handler fields
	// that the server uses the same network address and routes as before, and
	// the ErrorLog field so that the server now uses the custom errorLog logge
	// the event of any problems.
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(), // gets ServeMux from routes()
	}

	// Use the http.ListenAndServe() function to start a new web server. We pas
	// two parameters: the TCP network address to listen on (in this case ":4000
	// and the servemux we just created. If http.ListenAndServe() returns an er
	// we use the log.Fatal() function to log the error message and exit.
	// log.Printf("Starting server on %s\n", *addr)
	app.infoLog.Printf("Starting server on %s", *addr)
	//err := http.ListenAndServe(*addr, mux) // this goes to standard logger instead of error logger
	err = srv.ListenAndServe() // use struct error logger instead of standard logger

	// log.Fatal(err)
	errorLog.Fatal(err)
}

func openClient() (*firestore.Client, error) {
	// Use the application default credentials from environment variables
	client, dbErr := firestore.NewClient(context.Background(), "finfollow-app")
	if dbErr != nil {
		return nil, dbErr
	}

	// client.Collection("users").Doc("user1").Set(context.Background(), map[string]interface{}{
	// 	"first": "Ada",
	// 	"last":  "Lovelace",
	// 	"born":  1815,
	// })

	// client.Collection("users").Add(context.Background(), map[string]interface{}{
	// 	"first": "Julia",
	// 	"last":  "Sanchez",
	// 	"born":  2017,
	// })

	return client, nil
}
