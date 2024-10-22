package piscine

func AlphaCount(s string) int {
	dizi := []rune(s)
	sayac := 0
	for i := range dizi {
		if dizi[i] >= 'A' && dizi[i] <= 'Z' || dizi[i] >= 'a' && dizi[i] <= 'z' {
			sayac++
		}
	}
	return sayac
}
