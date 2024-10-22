package piscine

func BasicAtoi(s string) int {
	c := 0
	for _, i := range s {
		c = c*10 + int(i-'0')
	}
	return c
}
