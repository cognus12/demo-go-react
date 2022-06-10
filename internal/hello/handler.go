package hello

import (
	"log"
	"net/http"
)

func SayHello(w http.ResponseWriter, r *http.Request) {
	log.Println("Call api/hello")
	w.Write([]byte("Hello"))
}
