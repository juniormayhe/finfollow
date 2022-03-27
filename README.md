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

## Create Google service account

[Create IAM & Admin](https://cloud.google.com/docs/authentication/production#auth-cloud-implicit-go)

[Create service owner account](https://console.cloud.google.com/iam-admin/iam?project=finfollow-app&supportedpurview=project)

[Add JSON key to service account](https://console.cloud.google.com/iam-admin/serviceaccounts?project=finfollow-app&supportedpurview=project)

add path to environment variable
```
export GOOGLE_APPLICATION_CREDENTIALS="/path/to/jsonkey"

```