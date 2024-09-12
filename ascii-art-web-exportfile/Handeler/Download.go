package Handel

import (
	"net/http"
	"strconv"
)

func Download(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		Error(http.StatusMethodNotAllowed, w)
		return
	}
	//  mok  := os.ReadFile("css")
	// w.Header().Set("Content-Type", "text/css")
	// w.Write(mok)

	Asc := r.FormValue("Download")
	// fmt.Println(Asc)
	w.Header().Set("Content-Disposition", "attachment; filename=\"ascii-art.txt\"")
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Content-lent", strconv.Itoa(len(Asc)))
	w.Write([]byte(Asc))
}
