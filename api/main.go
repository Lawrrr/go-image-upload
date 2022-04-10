package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

func uploadImage(w http.ResponseWriter, r *http.Request) {
	// Set max image file size to 8mb
	MAX_UPLOAD_SIZE, err := strconv.Atoi(os.Getenv("MAX_FILE_SIZE"))
	if err != nil {
		fmt.Println("Value cannot be converted to integer")
		fmt.Println(err)
		return
	}
	AUTH_TOKEN := os.Getenv("AUTH_TOKEN")
	getToken := r.FormValue("auth")

	if AUTH_TOKEN == getToken {
		// Retrieve image from request
		file, handler, err := r.FormFile("data")
		if err != nil {
			fmt.Println("Error Retrieving the image")
			fmt.Println(err)
			return
		}
		defer file.Close()

		var contentType = handler.Header.Get("Content-Type")
		if strings.Contains(contentType, "image") && handler.Size < int64(MAX_UPLOAD_SIZE) {
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

			dst, err := ioutil.TempFile("public\\images", imgType)
			if err != nil {
				fmt.Println(err)
				return
			}
			defer dst.Close()

			// Store image info to database
			var dbName = os.Getenv("DB_NAME")
			db, err := sql.Open("sqlite3", "./"+dbName+".db")
			if err != nil {
				fmt.Println(err)
				return
			}
			defer db.Close()
			db.Exec("INSERT INTO images(path, content_type, size) values(?,?,?);", dst.Name(), contentType, handler.Size)

			// This is needed so that the image uploaded won't be corrupted
			fileBytes, err := ioutil.ReadAll(file)
			if err != nil {
				fmt.Println(err)
			}
			// Write this byte array to our uploaded image file
			dst.Write(fileBytes)

			fmt.Fprintf(w, "Image Uploaded!\n")
		} else {
			fmt.Fprintf(w, "File is not an image or image size is more than 8mb.")
		}
	} else {
		w.WriteHeader(403)
		fmt.Println("Forbidden: Incorrect auth key")
	}
}

func startEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func handleRequests() {
	var getBaseUrl = os.Getenv("BASE_URL")
	http.Handle("/", http.FileServer(http.Dir("./public/html")))
	http.HandleFunc("/upload", uploadImage)
	http.ListenAndServe(getBaseUrl+":8080", nil)
}

func main() {
	startEnv()
	handleRequests()
}
