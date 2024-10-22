package piscine

// ActiveBits, bir tam sayının ikili (binary) temsilindeki aktif (1) bitlerin sayısını hesaplar.
func ActiveBits(n int) (total int) {
	// n > 1 koşuluyla bir döngü başlatılır.
	// Döngü, n değerini her seferinde 2'ye böler ve her adımda en sağdaki biti kontrol eder.
	for ; n > 1; n /= 2 {
		// n'in en sağdaki biti (LSB) ile 1'i AND işlemi yaparak, bu bitin değerini kontrol eder.
		// Eğer bu bit 1 ise, total değişkenine 1 eklenir.
		total += n % 2
	}

	// Son olarak, n'in kendisi, yani son kalan bit kontrol edilir.
	// Eğer bu son bit 1 ise, total değişkenine 1 eklenir.
	total += n

	// Toplam aktif bit sayısı hesaplandıktan sonra, total değeri döndürülür.
	return total
}
