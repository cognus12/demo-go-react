package main

import (
	"html/template"
	"log"
	"net/http"
)

type ViewData struct {
	Title string
}

var index *template.Template

func loadTemplate() *template.Template {
	t, err := template.ParseFiles("frontend/dist/index.html")

	if err == nil {
		return t
	} else {
		panic(err)
	}
}

func init() {
	index = loadTemplate()
}

func handler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	data := ViewData{Title: "Go+Vite App"}

	index.Execute(w, data)
}

func handlerHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
}

func main() {
	fs := http.FileServer(http.Dir("frontend/dist/assets"))

	http.HandleFunc("/", handler)
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	http.HandleFunc("/api/hello", handlerHello)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
