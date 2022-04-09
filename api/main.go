package main

import (
	"fmt"
	"net/http"
)

func uploadImage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Image Uploaded!")
	fmt.Println("Endpoint Hit: upload image page")
}

func handleRequests() {
	http.HandleFunc("/upload", uploadImage)
	http.ListenAndServe(":8080", nil)
}

func main() {
	handleRequests()
}
