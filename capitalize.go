package piscine

func Capitalize(s string) string {
	dizi := []rune(s)
	for i, char := range dizi {
		if isnumberoralph(char) {
			if i == 0 || isnumberoralph(dizi[i-1]) == false {
				if dizi[i] >= 'a' && dizi[i] <= 'z' {
					dizi[i] = char - 32
				}
			} else {
				if dizi[i] >= 'A' && dizi[i] <= 'Z' {
					dizi[i] = char + 32
				}
			}
		}
	}

	return string(dizi)
}

func isnumberoralph(r rune) bool {
	if r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' || r >= '0' && r <= '9' {
		return true
	}
	return false
}
