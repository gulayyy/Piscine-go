package piscine

func NRune(s string, n int) rune {
	dizi := []rune(s)
	if n > len(dizi) {
		return 0
	} else {
		for a := 1; a <= len(dizi); a++ {
			if n == a {
				return dizi[n-1]
			}
		}
	}
	return 0
}
