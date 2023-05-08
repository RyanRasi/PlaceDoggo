package main

import (
	"bytes"
	"fmt"
	"image/jpeg"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/nfnt/resize"
)

// READ - All Items
func getAny(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/jpeg")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	img, err := os.Open("./assets/dog.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer img.Close()
	io.Copy(w, img)
}

// READ - Single Item
func getItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/jpeg")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	allParams := mux.Vars(r)["dimensions"]
	params := strings.Split(allParams, "x") // Get Params

	width, height := params[0], params[1]

	// open "test.jpg"
	file, err := os.Open("./assets/dog3.jpeg")
	if err != nil {
		log.Fatal(err)
	}

	// decode jpeg into image.Image
	img, err := jpeg.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	// resize to width 1000 using Lanczos resampling
	w64, err := strconv.ParseUint(width, 10, 32)
	if err != nil {
		fmt.Println(err)
	}
	wd := uint(w64)

	h64, err := strconv.ParseUint(height, 10, 32)
	if err != nil {
		fmt.Println(err)
	}
	ht := uint(h64)

	// and preserve aspect ratio
	newImage := resize.Resize(wd, ht, img, resize.Lanczos3)

	var opt jpeg.Options
	opt.Quality = 100

	buff := bytes.NewBuffer(nil)
	erro := jpeg.Encode(buff, newImage, &opt)
	if erro != nil {
		log.Fatal(erro)
	}

	io.Copy(w, buff)
}

func main() {
	// Server Port
	port := 8000

	// Initialise Router
	router := mux.NewRouter()

	// Router Handlers / Endpoints
	router.HandleFunc("/any", getAny).Methods("GET")
	router.HandleFunc("/get/{dimensions}", getItem).Methods("GET")

	fmt.Println("PlaceDoggo")
	fmt.Println("Web server started on port:", port)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), router))
}
