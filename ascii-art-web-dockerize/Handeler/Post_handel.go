package Handel

import (
	"html/template"
	"net/http"

	ascii "stylize/Ascii-art"
)

// handel the post request
func Post_handel(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" { // chaeck request methode
		Error(http.StatusMethodNotAllowed, w)
		return
	}
	if r.URL.Path != "/page.html" { // check the request path
		Error(http.StatusNotFound, w)
		return
	}
	banner := r.FormValue("banner")
	input := r.FormValue("text")

	if ascii.Checkbanner(banner) || !ascii.IsEmpty(input) || !ascii.Is_print(input) || len(input) > 1000 {
		Error(http.StatusBadRequest, w)
		return
	}
	template, err := template.ParseFiles("Templat/page.html")
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}

	cont := (ascii.Asci_art(input, banner))
	template.Execute(w, cont)
}
