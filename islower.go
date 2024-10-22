package piscine

func IsLower(s string) bool {
	dizi := []rune(s)
	ln := len(s)

	for i := 0; i <= ln-1; i++ {
		if dizi[i] < 'a' || dizi[i] > 'z' {
			return false
		}
	}
	return true
}
