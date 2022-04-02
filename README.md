# Run web and handlers

```
go run cmd/web/*
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

[Setup Firebase](https://firebase.google.com/docs/firestore/quickstart#go)
```
go get firebase.google.com/go
```

or update
```
go get -u firebase.google.com/go
```

Composable middleware
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

Session manager

Encrypted / authenticated cookie-based session store. Load and save via middleware. Stores up to 4KB.
```
go get github.com/golangcollege/sessions
```
We could use others like
- https://github.com/gorilla/sessions (better wait version 2 for memory leak fix and renewable session token)
- https://github.com/alexedwards/scs (load and save data via middleware)

## Create Google service account

[Create IAM & Admin](https://cloud.google.com/docs/authentication/production#auth-cloud-implicit-go)

[Create service owner account](https://console.cloud.google.com/iam-admin/iam?project=finfollow-app&supportedpurview=project)

[Add JSON key to service account](https://console.cloud.google.com/iam-admin/serviceaccounts?project=finfollow-app&supportedpurview=project)

add path to environment variable
```
export GOOGLE_APPLICATION_CREDENTIALS="/path/to/jsonkey"
export GOOGLE_APPLICATION_CREDENTIALS="/c/Users/junio/OneDrive/Outros Documentos/finfollow-app-6341c1702a75.json"
```

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