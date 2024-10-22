package piscine

func StrLen(s string) int {
	n := 0
	karakter := []rune(s)
	for a := range karakter {
		n = a
	}
	return n + 1
}
