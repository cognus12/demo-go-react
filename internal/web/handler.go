package web

import "net/http"

func HandleIndex(w http.ResponseWriter, r *http.Request) {

	index.Execute(w, data)
}

var fs = http.FileServer(http.Dir("static/assets"))
var HandleAssets = http.StripPrefix("/assets/", fs)
