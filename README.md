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