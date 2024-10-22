package piscine

func CountIf(f func(string) bool, tab []string) int {
	sayac := 0
	for _, i := range tab { // verilen fonksiyon içerisinde sayarak true yazdırıyor ve onların sayısını toplayıp yazıyor
		if f(i) == true {
			sayac++
		}
	}
	return sayac
}
