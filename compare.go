package piscine

func Compare(a, b string) int {
	if a == b { // alfabedeki indexlere göre sıraladık
		return 0
	} else if a < b {
		return -1
	} else {
		return 1
	}
}
