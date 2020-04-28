package filesystem

import (
	"net/http"
	"fmt"

	"io/ioutil"
)

//upload handler
func uploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("File Upload Endpoint Hit")

	// req contains 'multipartform'
	// req.ParseMultiPlatform parses the form
	// 10<<20 insicate max upload of 10MB
	r.ParseMultipartForm(10 << 20)

	// req. contains key 'myFile'.
	// req.FormFile returns/open file + handler (filename, header, size)
	file, handler, err := r.FormFile("myFile")
	if err != nil {
			fmt.Println("Error Retrieving the File")
			fmt.Println(err)
			return
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	// Create a temporary file
	tempFile, err := ioutil.TempFile("filesystem/temp-images", "upload-*.png")
	if err != nil {
			fmt.Println(err)
	}
	defer tempFile.Close()

	// read all of the file contents to []byte
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
			fmt.Println(err)
	}
	// write []byte to our temp file
	tempFile.Write(fileBytes)
	
	fmt.Fprintf(w, "Successfully Uploaded File\n")
}

func setupRoutes() {
	http.HandleFunc("/upload", uploadFile)
	http.ListenAndServe(":8080", nil)
}

func UploadFile() {
	setupRoutes()
	fmt.Println("hello World")
}