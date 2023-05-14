# PlaceDoggo
A placeholder image generator written in Go that returns specifically sized pictures of dogs.

### Endpoints

{localhost or IP}:8000/any - Returns a random picture from the assets folder 

{localhost or IP}:8000/get/{width}x{height} - Returns a random picture from the assets folder that has been resized based on your parameters

### How to Run

Install Go dependencies
```
go mod "github.com/gorilla/mux"
go mod "github.com/nfnt/resize"
```

Run with
```
go run main.go
```
