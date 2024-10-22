package piscine

import "github.com/01-edu/z01"

// PrintNum, verilen bir tamsayının rakamlarını ekrana yazdırmak için kullanılır.
// Bu fonksiyon, sayının her bir rakamını ayrı ayrı yazdırır.
func PrintNum(num int) {
	i := '0'
	if num == 0 {
		z01.PrintRune('0') // Eğer sayı 0 ise, sadece '0' karakterini yazdır.
		return
	}
	for j := 1; j <= num%10; j++ {
		i++ // Her basamağı '0' karakterinden başlayarak hesaplayarak rakamı bulur.
	}
	for j := -1; j >= num%10; j-- {
		i++ // Negatif sayılar için rakam hesaplamasını devam ettirir.
	}
	if num/10 != 0 {
		PrintNum(num / 10) // Sayının bir sonraki basamağını işlemek için kendisini tekrar çağırır.
	}
	z01.PrintRune(i) // Rakamı ekrana yazdırır.
	return
}

// PrintNbr, verilen bir tamsayıyı ekrana yazdırmak için kullanılır.
// Eğer sayı negatifse, negatif işareti yazdırır ve ardından rakamları yazdırmak için PrintNum fonksiyonunu çağırır.
func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune('-') // Eğer sayı negatifse, negatif işareti ("-") yazdır.
	}
	PrintNum(n) // Rakamları yazdırmak için PrintNum fonksiyonunu çağırır.
}
