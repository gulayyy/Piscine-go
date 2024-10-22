package piscine

func TrimAtoi(s string) int {
	nb := 0
	minus := false // negatif değer atama

	for _, v := range s {
		if v >= '0' && v <= '9' {
			digit := int(v - '0')
			nb = nb*10 + digit
		} else if nb == 0 && v == '-' {
			minus = true
		}
	}
	if minus {
		nb = -nb
	}
	return nb
}
