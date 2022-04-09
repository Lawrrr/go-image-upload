package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func uploadImage(w http.ResponseWriter, r *http.Request) {
	// Set max image file size to 8mb
	const MAX_UPLOAD_SIZE = 8388608

	// Retrieve image from request
	file, handler, err := r.FormFile("data")
	if err != nil {
		fmt.Println("Error Retrieving the image")
		fmt.Println(err)
		return
	}
	defer file.Close()

	var contentType = handler.Header.Get("Content-Type")
	if strings.Contains(contentType, "image") && handler.Size < MAX_UPLOAD_SIZE {
		// Create file and upload it to the images folder
		var imgType string
		switch contentType {
		case "image/jpeg":
			imgType = "upload-*.jpg"
			break
		case "image/gif":
			imgType = "upload-*.gif"
			break
		default:
			imgType = "upload-*.png"
			break
		}

		dst, err := ioutil.TempFile("public/images", imgType)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer dst.Close()

		// This is needed so that the image uploaded won't be corrupted
		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println(err)
		}
		// Write this byte array to our uploaded image file
		dst.Write(fileBytes)

		fmt.Fprintf(w, "Successfully Uploaded File\n")
	} else {
		fmt.Fprintf(w, "File is not an image or image size is more than 8mb.")
	}
}

func handleRequests() {
	http.Handle("/", http.FileServer(http.Dir("./public/html")))
	http.HandleFunc("/upload", uploadImage)
	http.ListenAndServe(":8080", nil)
}

func main() {
	handleRequests()
}
