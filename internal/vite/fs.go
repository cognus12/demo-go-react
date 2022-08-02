package vite

import (
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"net/url"
	"path"
)

type DevHandler struct{}

func copyHeader(dst, src http.Header) {
	for k, vv := range src {
		for _, v := range vv {
			dst.Add(k, v)
		}
	}
}

// TODO make cleaner
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

func (v *Vite) FileServer() http.Handler {
	stripPrefix := "/"

	if v.Env == "development" {
		return http.StripPrefix(stripPrefix, DevHandler{})
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		dist, err := fs.Sub(v.DistFS, path.Join(v.OutDir, v.AssetsDir))

		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		dirToServ := http.FS(dist)
		server := http.StripPrefix(v.AssetsURLPrefix, http.FileServer(dirToServ))

		server.ServeHTTP(w, r)
	}

	fshandler := http.HandlerFunc(handler)

	return fshandler
}
