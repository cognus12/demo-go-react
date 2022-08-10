package vite

import (
	"io/fs"
	"net/http"
	"path"
)

func (v *Vite) FileServer() http.Handler {
	if v.Env == "development" {
		return http.StripPrefix("/", &DevProxy{url: v.DevServerURL})
	}

	dist, err := fs.Sub(v.DistFS, path.Join(v.OutDir, v.AssetsDir))
	dirToServ := http.FS(dist)
	server := http.StripPrefix(v.AssetsURLPrefix, http.FileServer(dirToServ))

	handler := func(w http.ResponseWriter, r *http.Request) {
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		server.ServeHTTP(w, r)
	}

	fshandler := http.HandlerFunc(handler)

	return fshandler
}
