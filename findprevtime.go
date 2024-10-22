package piscine

// FindPrevPrime, verilen bir tamsayının (nb) öncesindeki en büyük asal sayıyı bulur.
func FindPrevPrime(nb int) int {
	if nb < 2 {
		return 0
	}
	if Isprime(nb) {
		return nb // Eğer nb bir asal sayıysa, nb'yi döndürür.
	}
	return FindPrevPrime(nb - 1) // Eğer nb bir asal sayı değilse, nb'yi bir azaltarak önceki sayıyı arar.
}

// Isprime, verilen bir tamsayının (n) asal olup olmadığını kontrol eder.
func Isprime(n int) bool {
	if n < 1 {
		return false // Negatif veya 0 olan sayılar asal değildir.
	} else if n == 2 || n == 3 {
		return true // 2 ve 3 asal sayılardır.
	} else if n%2 == 0 || n%3 == 0 {
		return false // 2 veya 3'e bölünebilen sayılar asal değildir.
	} else {
		// Diğer durumlar için asal sayı kontrolü yapılır.
		// 6k ± 1 formülü kullanılarak asal sayıların kontrolü yapılır.
		for i := 5; i*i <= n; i += 6 {
			if n%i == 0 || n%(i+2) == 0 {
				return false // Sayı i veya i+2 ile tam bölünebiliyorsa, asal değildir.
			}
		}
	}
	return true // Yukarıdaki koşullardan hiçbiri sağlanmıyorsa, sayı asal olarak kabul edilir.
}
