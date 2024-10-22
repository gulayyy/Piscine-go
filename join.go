package piscine

func Join(strs []string, sep string) string {
	str := ""
	for a, i := range strs {
		str += i
		if !(a == len(strs)-1) {
			str += sep
		}
	}
	return str
}
