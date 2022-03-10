package server

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var templates = template.Must(template.ParseFiles("client/index.html"))

func HomePage(w http.ResponseWriter, r *http.Request) {
	url := struct {
		URL string
	}{os.Getenv("LULUPOD_URL")}
	templates.ExecuteTemplate(w, "index.html", url)
}

func UploadFile(w http.ResponseWriter, r *http.Request) {
	// maximum upload of about 6.4 GB
	r.ParseMultipartForm(6 << 30)

	file, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	fmt.Printf("***\nUploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n***\n", handler.Size)

	dst, err := os.Create("./public/" + handler.Filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	HomePage(w, r)
}

func ReadFiles(w http.ResponseWriter, r *http.Request) {
	var files []string
	err := filepath.Walk("./public", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}
		files = append(files, filepath.Base(path))
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	res, err := json.Marshal(files[1:])
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(res)
}
