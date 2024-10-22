package piscine

func IsNumeric(s string) bool {
	dizi := []rune(s)
	ln := len(s)

	for i := 0; i <= ln-1; i++ {
		if dizi[i] < '0' || dizi[i] > '9' {
			return false
		}
	}
	return true
}
