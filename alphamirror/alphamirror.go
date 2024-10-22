package main

import (
	"fmt"
	"os"
)

func main() {
	// Eğer komut satırında 2 argüman varsa devam et
	if len(os.Args) == 2 {
		// Argümanı rune dizisine dönüştür
		bos := []rune(os.Args[1])
		// Dizide dolaş ve her bir karakter için işlem yap
		for i, char := range bos {
			// Eğer karakter küçük harfse
			if char >= 'a' && char <= 'z' {
				// Küçük harfi tersine çevir ve bos dizisine ata
				bos[i] = 'z' - char + 'a'
				// Eğer karakter büyük harfse
			} else if char >= 'A' && char <= 'Z' {
				// Büyük harfi tersine çevir ve bos dizisine ata
				bos[i] = 'Z' - char + 'A'
			}
		}
		// Tersine çevrilmiş stringi yazdır
		fmt.Println(string(bos))
	}
	// Bir satır boşluk bırak
	fmt.Println('\n')
}

/*bos[i] = 'z' - char + 'a' burada char değerini çıkarmamızın sebebi char orda bizim tersine çvireceğimiz karakter ve bu karakteri z değerinin ascii kodundan çıkarıp a degerinin
ascii koduna ekliyoruz ve böylece aranan ters değeri buluyoruz*/
