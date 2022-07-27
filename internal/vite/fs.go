package vite

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type DevHandler struct{}

func copyHeader(dst, src http.Header) {
	for k, vv := range src {
		for _, v := range vv {
			dst.Add(k, v)
		}
	}
}

func (DevHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	client := &http.Client{}
	r.RequestURI = ""

	var err error
	r.URL, err = url.Parse("http://localhost:3000/" + r.URL.Path)

	if err != nil {
		fmt.Println(err)
	}

	resp, err := client.Do(r)

	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	copyHeader(w.Header(), resp.Header)
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}

// func devHandler(w http.ResponseWriter, r *http.Request) {
// 	http.Redirect(w, r, "http://localhost:3000"+r.URL.Path, 300)
// }

func (v *Vite) FileServer() http.Handler {

	if v.Env == "development" {
		return http.StripPrefix("/", DevHandler{})
	}

	fs := http.FileServer(http.Dir(v.AssetsPath))

	return http.StripPrefix(v.AssetsURLPrefix, fs)
}
