package ascii

import (
	"strings"
)

// ascii-art normal
func Asci_art(inputtext string,banner string) string {
	// content := strings.Replace(inputtext, "\r\n", "\\n", -1)
	inputsep := strings.Split(inputtext, "\r\n")
	// if strings.ReplaceAll(inputtext, "\\n", "") == "" && len(inputtext) > 1 {
	// 	inputsep = inputsep[1:]
	// }
	contentfile := Readfile(banner)

	matrix := Split((string(contentfile)), banner)
	return Print(inputsep, matrix)
}
