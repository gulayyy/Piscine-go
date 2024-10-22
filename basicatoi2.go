package piscine

func BasicAtoi2(s string) int {
	c := 0
	for _, i := range s {
		if i < '0' || i > '9' {
			return 0
		} else {
			c = c*10 + int(i-'0')
		}
	}
	return c
}
