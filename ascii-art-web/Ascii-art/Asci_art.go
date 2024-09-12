package ascii

import (
	"strings"
)

// ascii-art normal
func Asci_art(inputtext string, banner string) string {
	inputsep := strings.Split(inputtext, "\r\n")
	contentfile := Readfile(banner)
	matrix := Split((string(contentfile)), banner)
	return Print(inputsep, matrix)
}
