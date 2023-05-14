package main

import (
	"bytes"
	"fmt"
	"image/jpeg"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/nfnt/resize"
)

// READ - All Items
func getAny(w http.ResponseWriter, r *http.Request) {
	// Set Headers
	w.Header().Set("Content-Type", "image/jpeg")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Call random int function
	picID := generateRandomInt()

	// Open specfic image based on int
	img, err := os.Open("./assets/" + strconv.Itoa(picID) + ".jpeg")
	if err != nil {
		log.Fatal(err)
	}
	defer img.Close()

	// Display image onto webpage
	io.Copy(w, img)
}

// READ - Single Item
func getItem(w http.ResponseWriter, r *http.Request) {
	// Set Headers
	w.Header().Set("Content-Type", "image/jpeg")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Read params
	allParams := mux.Vars(r)["dimensions"]
	params := strings.Split(allParams, "x") // Get Params

	width, height := params[0], params[1]

	// Call random int function
	picID := generateRandomInt()

	// Open specfic image based on int
	file, err := os.Open("./assets/" + strconv.Itoa(picID) + ".jpeg")
	if err != nil {
		log.Fatal(err)
	}

	// Decode jpeg into image.Image
	img, err := jpeg.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	// Convert params to unsigned ints
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

	// resize to width based on params using Lanczos resampling
	newImage := resize.Resize(wd, ht, img, resize.Lanczos3)

	var opt jpeg.Options
	opt.Quality = 100

	// Write image to memory
	buff := bytes.NewBuffer(nil)
	erro := jpeg.Encode(buff, newImage, &opt)
	if erro != nil {
		log.Fatal(erro)
	}

	// Display converted image onto webpage
	io.Copy(w, buff)
}

func generateRandomInt() int {

	// Folder path
	folderPath := "./assets/"

	// Read files in folder path
	files, err := ioutil.ReadDir(folderPath)
	if err != nil {
		fmt.Println(err)
		return 0
	}

	// Count all files in path
	count := 0
	for _, file := range files {
		if !file.IsDir() {
			count++
		}
	}
	count--

	// Generate number based on time seed
	randSeed := rand.New(rand.NewSource(time.Now().UnixNano()))
	number := randSeed.Intn(count) + 1

	// Return the generated number
	return number
}

func main() {
	// Server Port
	port := 8000

	// Initialise Router
	router := mux.NewRouter()

	// Router Handlers / Endpoints
	router.HandleFunc("/any", getAny).Methods("GET")
	router.HandleFunc("/get/{dimensions}", getItem).Methods("GET")

	// Listener and logs
	fmt.Println("PlaceDoggo")
	fmt.Println("Web server started on port:", port)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), router))
}
