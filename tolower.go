package piscine

func ToLower(s string) string {
	dizi := []rune(s)
	ln := len(s)

	for i := 0; i <= ln-1; i++ {
		if dizi[i] >= 'A' && dizi[i] <= 'Z' { // Büyük harfleri küçük harfe çevirmek için +32 ekliyoruz
			dizi[i] += 32
		}
	}
	return string(dizi)
}
