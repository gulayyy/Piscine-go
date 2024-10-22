package piscine

func Any(f func(string) bool, a []string) bool {
	for _, i := range a {
		if f(i) == true { // Verilen fonsksiyon içinde dolaşıyor ve döndürülen değerin boolean değerine göre sonuç çıkarıyor
			return true
		}
	}
	return false
}
