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
	fmt.Println("Server is running at http://localhost:8081")
	http.ListenAndServe(":8081", nil)
}
