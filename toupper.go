package piscine

func ToUpper(s string) string {
	dizi := []rune(s)
	ln := len(s)

	for i := 0; i <= ln-1; i++ {
		if dizi[i] >= 'a' && dizi[i] <= 'z' {
			dizi[i] -= 32 // Küçük olan harfleri büyük olan harfe çevirmek için -32 çıkarıyoruz
		}
	}
	return string(dizi)
}
