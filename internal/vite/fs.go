package vite

import "net/http"

func (v *Vite) FileServer() http.Handler {
	fs := http.FileServer(http.Dir(v.AssetsPath))

	return http.StripPrefix(v.AssetsURLPrefix, fs)
}
