package piscine

import (
	"os"
	"strconv"
)

// GetNumber, komut satırından alınan argümanları işleyerek iki tamsayı ve bir başarı durumu döndüren bir fonksiyondur.
func GetNumber() (x, y int, ok bool) {
	// Komut satırı argümanları os.Args dizisinden alınır.
	args := os.Args[1:]

	// Eğer argüman sayısı 2 değilse, hemen hata durumu döndürülür.
	if len(args) != 2 {
		return 0, 0, false
	}

	// İki argüman, strconv.Atoi fonksiyonu kullanılarak tam sayılara dönüştürülür.
	x, hata1 := strconv.Atoi(args[0]) //"strconv" dizeyi (string) tam sayıya (int) dönüştürmek için kullanılır
	y, hata2 := strconv.Atoi(args[1])

	// Dönüşüm sırasında hata varsa (err1 veya err2 nil değilse), hemen hata durumu döndürülür.
	if hata1 != nil || hata2 != nil { // nil (hata olmadığını temsil eden değer)
		return x, y, false
	}

	// Eğer her iki tam sayı da negatifse, yine hata durumu döndürülür.
	if x < 0 || y < 0 {
		return x, y, false
	}

	// Yukarıdaki tüm koşullar sağlanmışsa, işlem başarılıdır ve true ile işaretleme yapılır.
	return x, y, true
}
