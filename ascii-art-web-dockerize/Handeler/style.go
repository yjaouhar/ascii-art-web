package Handel

import (
	"net/http"
	"strings"
)

func Style(w http.ResponseWriter, r *http.Request) {
	style := http.StripPrefix("/css/", http.FileServer(http.Dir("css")))
	if strings.HasSuffix(r.URL.Path, "/") {
		Error(http.StatusNotFound, w)
		return
	}
	style.ServeHTTP(w, r)
}
