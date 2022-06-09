package main

import (
	"govite/internal/web"
	"html/template"
	"log"
	"net/http"
	"os"
)

var mode string

var index *template.Template
var data *web.ViewData

func init() {
	mode = os.Getenv("MODE")

	if mode == "api" {
		return
	}

	index = web.LoadTemplate()
	data = web.PrepareIndexPageData()
}

func handler(w http.ResponseWriter, r *http.Request) {
	index.Execute(w, data)
}

func handlerHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
}

func main() {
	log.Println("Run server in mode: ", mode)

	if mode == "full" {
		fs := http.FileServer(http.Dir("static/assets"))
		http.HandleFunc("/", handler)
		http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	}

	http.HandleFunc("/api/hello", handlerHello)

	log.Println("Start server on localhost:8000")
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
