package main

import (
	"net/http"
	"path"
)

func (hs *HttpServer) isAccessibleResource(p string) bool {
	if restricted, _ := path.Match("/al[09]_", p); !restricted {
		return true
	}

	return false
}

func (hs *HttpServer) resourcesHandler(w http.ResponseWriter, r *http.Request) {

	p := r.URL.Path[1:] //strip '/', TODO: rteki, maybe there's some more elegan way

	if r.Method == http.MethodGet && hs.isAccessibleResource(p) {
		http.ServeFile(w, r, p)
	}

}
