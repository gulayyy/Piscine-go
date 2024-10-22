package main

import (
	"fmt"
)

// Bol fonksiyonu, bir dizeyi belirli bir ayıracı kullanarak böler ve sonuç olarak bir dize dilimi (slice) döndürür.
func Split(s, sep string) []string {
	var sonuc []string        // Sonuç dilimi (slice) için boş bir dize dilimi oluşturulur.
	baslangicIndeksi := 0     // Başlangıç indeksi, dizedeki parçanın başlangıç noktasını belirlemek için kullanılır.
	ayiracUzunlugu := len(sep) // Ayıracın uzunluğu, dizede aranan ayıracı belirlemek için kullanılır.

	// Döngü, dizedeki her bir karakteri kontrol eder ve ayıracı bulmaya çalışır.
	for i := 0; i <= len(s)-ayiracUzunlugu; i++ {
		// Dize içinde belirlenen konumda, belirlenen ayıraç ile eşleşme kontrol edilir.
		if s[i:i+ayiracUzunlugu] == sep {
			// Eğer ayırac bulunursa, parçanın başlangıç indeksinden bulunan indekse kadar olan kısmı sonuç dilimine eklenir.
			sonuc = append(sonuc, s[baslangicIndeksi:i])
			// Başlangıç indeksi, bulunan ayıracın sonrasındaki kısma kaydırılır.
			baslangicIndeksi = i + ayiracUzunlugu
		}
	}

	// Son olarak, işlenmemiş kalan kısmı sonuç dilimine eklemek için bir kontrol yapılır.
	if baslangicIndeksi < len(s) {
		sonuc = append(sonuc, s[baslangicIndeksi:])
	}

	// Sonuç dilimi döndürülür.
	return sonuc
}


func main() {
	s := "HelloHAhowHAareHAyou?"
	fmt.Printf("%#v\n", Split(s, "HA"))
}
