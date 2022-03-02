package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/noam-g4/lulupod/server"
)

func main() {

	fmt.Printf("===========\nLULUPOD 1.0\n===========\n")

	fs := http.FileServer(http.Dir("./public"))

	http.HandleFunc("/", server.HomePage)
	http.HandleFunc("/list", server.ReadFiles)
	http.HandleFunc("/upload", server.UploadFile)
	http.Handle("/files/", http.StripPrefix("/files", fs))

	err := http.ListenAndServe(":8000", nil)
	log.Fatal(err)

}
