package piscine

func IsSorted(f func(int, int) int, arr []int) bool {
	a := true
	b := true

	for i := 1; i < len(arr); i++ { // artana göre sıralama yapıyor
		if !(f(arr[i-1], arr[i]) >= 0) { // verilen f fonksiyonu içerisine  a<b a=b a>b şeklinde kontrol edip ona göre kıyaslama yapıyoruz
			a = false
		}
	}
	for i := 1; i < len(arr); i++ { // azalana göre sırlama yapıyor
		if !(f(arr[i-1], arr[i]) <= 0) {
			b = false
		}
	}
	return a || b // birininin true döndürmesi durumunda da  true döndürüyor
}
