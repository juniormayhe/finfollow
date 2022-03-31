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