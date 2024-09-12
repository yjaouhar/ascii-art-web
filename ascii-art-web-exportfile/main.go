package main

import (
	"fmt"
	"net/http"

	Handel "stylize/Handeler"
)

func main() {
	http.HandleFunc("/css/", Handel.Style)
	http.HandleFunc("/", Handel.Get_handel)
	http.HandleFunc("/page.html", Handel.Post_handel)
	http.HandleFunc("/Download", Handel.Download)
	fmt.Println("Server is running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
