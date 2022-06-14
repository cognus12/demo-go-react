package main

import (
	"demo-go-react/internal/hello"
	"demo-go-react/internal/web"
	"log"
	"net/http"
	"os"
)

func main() {
	mode := os.Getenv("MODE")

	log.Printf("Run app in %v mode \n", mode)

	if mode == "full" {
		http.HandleFunc("/", web.HandleIndex)
		http.Handle("/assets/", web.HandleAssets)

		log.Println("Static assets loaded")
	}

	http.HandleFunc("/api/hello", hello.SayHello)

	log.Println("Start server on localhost:8000")

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
