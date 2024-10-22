package piscine

func IsAlpha(s string) bool {
	dizi := []rune(s)
	ln := len(s)

	for i := 0; i <= ln-1; i++ {
		if (dizi[i] < 'a' || dizi[i] > 'z') && (dizi[i] < 'A' || dizi[i] > 'Z') && (dizi[i] < '0' || dizi[i] > '9') {
			return false
		}
	}
	return true
}
