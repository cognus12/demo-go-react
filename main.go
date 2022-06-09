package main

import (
	"govite/internal/manifest"
	"html/template"
	"log"
	"net/http"
	"os"
)

type ViewData struct {
	Title string
	// Favicon string
	*manifest.Chunck
}

var index *template.Template
var resources *manifest.ManifestData
var mainChunck *manifest.Chunck

func loadTemplate() *template.Template {
	t, err := template.ParseFiles("template.html")

	if err == nil {
		return t
	} else {
		panic(err)
	}
}

func init() {
	var manifestError error

	resources, manifestError = manifest.Parse("static/manifest.json")

	if manifestError != nil {
		log.Fatal("Failed to load manifest.json")
	}

	mainChunck = manifest.GetMain(resources)

	if mainChunck == nil {
		log.Fatal("Failed to load assets")
	}

	index = loadTemplate()
}

func handler(w http.ResponseWriter, r *http.Request) {
	data := ViewData{Title: "Go+Vite App", Chunck: *&mainChunck}

	index.Execute(w, data)
}

func handlerHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
}

func main() {
	fs := http.FileServer(http.Dir("static/assets"))

	// TODO implement running in modes - full, api
	mode := os.Getenv("MODE")

	log.Println("Run server in mode: ", mode)

	http.HandleFunc("/", handler)
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	http.HandleFunc("/api/hello", handlerHello)

	log.Println("Start server on localhost:8000")
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
