package piscine

func LastRune(s string) rune {
	dizi := []rune(s)
	return dizi[len(s)-1]
}
