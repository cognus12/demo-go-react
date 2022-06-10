package main

import (
	"govite/internal/web"
	"log"
	"net/http"
	"os"
)

func handlerHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
}

func main() {
	mode := os.Getenv("MODE")

	if mode == "full" {
		http.HandleFunc("/", web.HandleIndex)
		http.Handle("/assets/", web.HandleAssets)
	}

	http.HandleFunc("/api/hello", handlerHello)

	log.Println("Start server on localhost:8000")
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
