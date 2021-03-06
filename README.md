## Create Google service account

[Create IAM & Admin](https://cloud.google.com/docs/authentication/production#auth-cloud-implicit-go)

[Create service owner account](https://console.cloud.google.com/iam-admin/iam?project=finfollow-app&supportedpurview=project)

[Add JSON key to service account](https://console.cloud.google.com/iam-admin/serviceaccounts?project=finfollow-app&supportedpurview=project)

add path to environment variable
```
export GOOGLE_APPLICATION_CREDENTIALS="/path/to/jsonkey"
export GOOGLE_APPLICATION_CREDENTIALS="/c/Users/junio/OneDrive/Outros Documentos/finfollow-app-6341c1702a75.json"
```

# Run web and handlers

```
go run cmd/web/*
```

Run web server without excluding test files, enable globbling in bash
```
shopt -s extglob
go run cmd/web/!(*_test).go
```

# Testing

Run tests with verbose mode
```
go test -v cmd/web/*
```

Stop test running after first failure
```
go test -failfast -v cmd/web/*
```

Run a specific test
```
go test -v -run="^TestPing$" ./cmd/web
```

Run a specific test with subtests
```
go test -v -run="^TestHumanDate$/^UTC|CET$" ./cmd/web
```

Run all tests starting from current folder
```
go test ./...
```

Run tests in cmd/web folder
```
go test ./cmd/web/
```

Run maximum of 4 tests in parallel in all folders
```
go test -parallel 4 ./...
```

## Use arguments
```
go run cmd/web/* -help
go run cmd/web/* -addr=":9999"
```

## Redirect stdout and stderr to append outputs to disk files
```
go run cmd/web/* >>/tmp/info.log 2>>/tmp/error.log
```

# Dependencies

## Database
[Setup Firebase](https://firebase.google.com/docs/firestore/quickstart#go)
```
go get firebase.google.com/go
```

or update
```
go get -u firebase.google.com/go
```

## Composable middleware
```
go get github.com/justinas/alice
```

Router to deal with method based routing and clean urls, not maintained. We could use others like
- https://github.com/go-zoo/bone
- https://github.com/go-chi/chi
- https://github.com/gorilla/mux
- 
```
go get github.com/bmizerany/pat
```

## Session manager

Encrypted / authenticated cookie-based session store. Load and save via middleware. Stores up to 4KB.
```
go get github.com/golangcollege/sessions
```
We could use others like
- https://github.com/gorilla/sessions (better wait version 2 for memory leak fix and renewable session token)
- https://github.com/alexedwards/scs (load and save data via middleware)

## Password handling

Saving passwords with bcrypt implementation, available in golang.org/x/crypto package installed along with golangcollege/sessions.

4096 (2^12) bcrypt iterations will be used to hash the password, returning a 60-character long hash.

## CSRF protection

Since SameSite=Strict may not work in all browsers, we
can use a 3rd-party package to generate a random a CSRF token as cookie and added to a hidden field.
```
go get github.com/justinas/nosurf
```

Alternative
- https://github.com/gorilla/csrf

## Methods
```
curl -iL -X POST http://localhost:4000/asset/add 
```

TODO:
```
curl -iL -X POST http://localhost:4000/asset/add -H "Content-Type: application/json" -d '{"name":"test", "value": 50.4, "currency":  "USD", "custody": "FIAT", created: "2020-01-01", active: true}'
```

## Self signed certificate  
In ./tls folder run
```
go run '/c/Program Files/Go/src/crypto/tls/generate_cert.go' --rsa-bits=2048 --host=localhost
```
