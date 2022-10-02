# workout-api
workout api with workout names and gifs

## Downloading workout gifs
Parses csv, downloads gifs
```bash
$ go run ./cmd/dlgifs/main.go
```

## seed workout db
Parses csv, seeds DB
```bash
$ go run ./cmd/seed/main.go
```

# Build 

```bash
$ go build main.go # serves on port 8080
```
