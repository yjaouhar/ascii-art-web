package Handel

import (
	"fmt"
	"html/template"
	"net/http"
)

// handel the get methode request
func Get_handel(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" { // check the request methode
		Error(http.StatusMethodNotAllowed, w)
		return
	}
	if r.URL.Path != "/" { // check the request path
		Error(http.StatusNotFound, w)
		return
	}

	template, err := template.ParseFiles("Templat/page.html") // parse file page.html
	if err != nil {
		fmt.Println("mok")
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}
	template.Execute(w, nil)
}
