package Handel

import (
	"html/template"
	"net/http"
)

// handel the get methode request
func Get_handel(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {  // check the request methode
		http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path != "/" { // check the request path
		http.Error(w, "404 Page Not Found", http.StatusNotFound)
		return
	}
	template, err := template.ParseFiles("../page web/page.html") // parse file page.html 
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}
	template.Execute(w, nil)
}
