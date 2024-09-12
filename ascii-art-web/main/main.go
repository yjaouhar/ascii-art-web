package main

import (
	"fmt"
	"net/http"

	"ascii-art-web/Handel"
)

func main() {
	http.HandleFunc("/", Handel.Get_handel)
	http.HandleFunc("/page.html", Handel.Post_handel)
	fmt.Println("server run in http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
