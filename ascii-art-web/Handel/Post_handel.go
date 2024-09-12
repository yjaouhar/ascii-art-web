package Handel

import (
	"html/template"
	"net/http"

	ascii "ascii-art-web/Ascii-art"
)

// handel the post request 
func Post_handel(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" { // chaeck request methode 
		http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path != "/page.html" { // check the request path
		http.Error(w, "404 Page Not Found", http.StatusNotFound)
		return
	}
	banner := r.FormValue("banner")
	input := r.FormValue("text")

	if ascii.Checkbanner(banner) || !ascii.IsEmpty(input) || !ascii.Is_print(input) || len(input) > 1000 {

		http.Error(w, "400 Bad Request", http.StatusBadRequest)
		return
	}
	template, err := template.ParseFiles("../page web/page.html")
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}
	
	cont := (ascii.Asci_art(input, banner))
	template.Execute(w, cont)
}
