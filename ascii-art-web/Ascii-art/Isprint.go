package ascii

// check the  argument is have a caracter printible
func Is_print(s string) bool {
	slr := []rune(s)
	print := true
	for i := 0; i < len(slr); i++ {
		if !(slr[i] >= 10 && slr[i] <= 126) {
			print = false
		}
	}
	return print
}
