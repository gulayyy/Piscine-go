package piscine

func IsUpper(s string) bool {
	dizi := []rune(s)
	ln := len(s)

	for i := 0; i <= ln-1; i++ {
		if dizi[i] < 'A' || dizi[i] > 'Z' {
			return false
		}
	}
	return true
}
