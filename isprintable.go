package piscine

func IsPrintable(s string) bool {
	dizi := []rune(s)
	ln := len(s)
	for i := 0; i <= ln-1; i++ {
		if dizi[i] < ' ' {
			return false
		}
	}
	return true
}
