package main

import (
	"govite/internal/manifest"
	"html/template"
	"log"
	"net/http"
)

type ViewData struct {
	Title string
	// Favicon string
	Js string
}

var index *template.Template

func loadTemplate() *template.Template {
	t, err := template.ParseFiles("template.html")

	if err == nil {
		return t
	} else {
		panic(err)
	}
}

func init() {
	resources, err := manifest.Parse("static/manifest.json")

	if err != nil {
		log.Fatal("Failed to load manifest.json")
	}

	log.Printf("Loaded manifest data: %v \n", resources)

	index = loadTemplate()
}

func handler(w http.ResponseWriter, r *http.Request) {
	data := ViewData{Title: "Go+Vite App"}

	index.Execute(w, data)
}

func handlerHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
}

func main() {
	fs := http.FileServer(http.Dir("static/assets"))

	http.HandleFunc("/", handler)
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	http.HandleFunc("/api/hello", handlerHello)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
