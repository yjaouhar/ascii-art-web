package Handel

import (
	"html/template"
	"net/http"
)

type messag struct {
	Status  int
	Message string
}

// hndel error
func Error(status int, w http.ResponseWriter) {
	temp, err := template.ParseFiles("Templat/Error.html")
	if err != nil {
		http.Error(w, "500 enternal server error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(status)
	errore := messag{
		Status:  status,
		Message: http.StatusText(status),
	}

	temp.Execute(w, errore)
}
