package web

import (
	"log"
	"net/http"
)

func ServeIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		log.Println("Execute index.html")

		index.Execute(w, data)
	}
}

var fs = http.FileServer(http.Dir("static/assets"))
var HandleAssets = http.StripPrefix("/assets/", fs)
