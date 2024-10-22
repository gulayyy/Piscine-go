package piscine

func sirala(sayilar []int) {
	for i := 0; i < len(sayilar); i++ {
		for j := 0; j < len(sayilar); j++ {
			if sayilar[j] > sayilar[i] {
				sayilar[i], sayilar[j] = sayilar[j], sayilar[i]
			}
		}
	}
}

func Abort(a, b, c, d, e int) int {
	sayilar := []int{a, b, c, d, e}
	sirala(sayilar)
	return sayilar[2]
}
