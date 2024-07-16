package main

import (
	"archive/zip"
	"errors"
	"log"
	"net/http"
	"os"
)

func main() {
	filename := os.Args[1];
	if len(os.Args) < 2 {
		log.Fatal("first argument must be a zipfile")
	}
	log.Fatal(run(filename))
}

func run(filename string) error {
	
	if filename == "" {
		return errors.New("need a path to a zip file")
	}
	zr, err := zip.OpenReader(filename)
	if err != nil {
		return err
	}
	defer zr.Close()
	const addr = "localhost:8000"
	log.Printf("see http://%s/", addr)
	return http.ListenAndServe(addr, http.FileServer(http.FS(zr)))
}
