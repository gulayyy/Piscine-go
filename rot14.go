package piscine

func Rot14(s string) string {
	dizi := []rune(s)
	for i := 0; i < len(dizi); i++ {
		if dizi[i] >= 'A' && dizi[i] <= 'L' || dizi[i] >= 'a' && dizi[i] <= 'l' {
			dizi[i] = dizi[i] + 14
		} else if dizi[i] >= 'M' && dizi[i] <= 'Z' || dizi[i] >= 'm' && dizi[i] <= 'z' {
			dizi[i] = dizi[i] - 12
		}
	}
	return string(dizi)
}
