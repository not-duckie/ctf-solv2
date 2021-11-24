package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"text/template"
)

var templates = template.Must(template.ParseFiles("public/upload.html"))

func display(w http.ResponseWriter, page string, data interface{}) {
	templates.ExecuteTemplate(w, page+".html", data)
}

func uploadFile(w http.ResponseWriter, r *http.Request) {

	// Get handler for filename, size and headers
	file, handler, err := r.FormFile("File")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error Retrieving the File"))
		log.Println(err)
		return
	}

	defer file.Close()
	data, _ := io.ReadAll(file)
	if checkVera(data) {

		//writing data to file
		f, err := os.Create(".\\uploads\\" + handler.Filename)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		f.Write(data)

		fmt.Fprintf(w, "Successfully Uploaded File\n")
	} else {
		fmt.Fprintf(w, "Not a Vera file")
	}
}

func checkVera(data []byte) bool {
	/*
		Since vera encrypted virtual disks are not identifiable thus following 2 layered
		approach to this.

		1) we check if the file does not have have matching header to known format.
		2) Second the entropy of the data should be greater than 7.5 as it indicates encrypted data.

	*/
	if entropy(data) > 7.5 {
		if checkFileType(data) {
			return true
		}
	}
	return false
}
