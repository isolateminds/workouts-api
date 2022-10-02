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

# Usage 

```
 GET http://localhost:8080/api/workouts?search=tricep+dips
```

## Result
```json
[
    {
        "bodypart": "upper arms",
        "equipment": "weighted",
        "gifurl": "/api/gifs/1755.gif",
        "id": "1755",
        "name": "weighted tricep dips",
        "target": "triceps"
    }
]
```
## Getting gifs 
```
GET http://localhost:8080/api/gifs/1755.gif
```

## Result

![1755](https://user-images.githubusercontent.com/112124260/193437052-ed48b30f-92ec-4a0f-ae24-3e7373f88e25.gif)


