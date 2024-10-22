package main

// Itoa, bir tamsayıyı dize olarak temsil eder.
// Örneğin, 123 sayısı "123" dizesine dönüştürülür.
func Itoa(a int) string {
	var buf []byte // Tamsayının basamaklarını depolamak için bir byte dizisi.
	var r []byte   // Sonuç dizesini tutacak byte dizisi.
	var next int   // Bir sonraki basamağı hesaplamak için kullanılan değişken.
	var right int  // Tamsayının sağ tarafındaki basamağı tutacak değişken.

	// Tamsayının tüm basamaklarını alana kadar döngü devam eder.
	for {
		if a < 0 {
			a = -1 * a
			r = append(r, '-') // Eğer tamsayı negatifse, işareti sonuç dizisine ekler.
		}
		next = a / 10         // Bir sonraki basamağı hesaplar.
		right = a - next*10   // Sağ taraftaki basamağı alır.
		a = next              // Bir sonraki basamağa geçer.
		buf = append(buf, byte('0'+right)) // Alınan basamağı buf dizisine ekler.
		if a == 0 {
			break // Tüm basamaklar alındığında döngüden çıkar.
		}
	}

	// Alınan basamaklar, terslenerek sonuç dizisine eklenir.
	for j := 0; j < len(buf); j++ {
		r = append(r, buf[len(buf)-j-1])
	}

	// Sonuç dizesi byte dizisinden bir stringe dönüştürülerek döndürülür.
	return string(r)
}
