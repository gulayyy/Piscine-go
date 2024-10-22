package piscine

// IsPrime1, bir tamsayının (a) asal bir sayı olup olmadığını kontrol eder.
func IsPrime1(a int) bool {
	if a <= 1 {
		return false // Eğer a 1'den küçük veya negatifse, asal değildir.
	} else if a == 2 {
		return true // Eğer a 2 ise, asaldır.
	} else if a%2 == 0 {
		return false // Eğer a 2'ye tam bölünüyorsa, asal değildir.
	}
	for i := 3; i*i <= a; i += 2 {
		if a%i == 0 {
			return false // Diğer durumlar için asal sayı kontrolü yapılır ve eğer bölen bulunursa asal değildir.
		}
	}
	return true // Hiçbir bölen bulunmazsa, a asal bir sayıdır.
}

// FindNextPrime, verilen bir tamsayının (nb) sonraki asal sayısını bulur.
func FindNextPrime(nb int) int {
	for i := nb; i <= nb*2; i++ {
		if IsPrime1(i) {
			return i // Verilen sayıdan başlayarak sonraki asal sayıyı bulur ve döndürür.
		}
	}
	return 2 // Eğer sonraki asal sayı bulunamazsa, varsayılan olarak 2'yi döndürür.
}
