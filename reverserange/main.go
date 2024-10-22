package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		return // İki argüman verilmemişse veya yanlış giriş yapılmışsa programdan çıkar.
	}
	a, err := strconv.Atoi(os.Args[1]) // İlk komut satırı argümanını tamsayıya dönüştürür.
	if err != nil {
		fmt.Println(err) // Dönüştürme hatası varsa hatayı ekrana yazdırır.
		return           // Programı sonlandırır.
	}
	b, err := strconv.Atoi(os.Args[2]) // İkinci komut satırı argümanını tamsayıya dönüştürür.
	if err != nil {
		fmt.Println(err) // Dönüştürme hatası varsa hatayı ekrana yazdırır.
		return           // Programı sonlandırır.
	}
	fmt.Print(b) // İkinci tamsayıyı ekrana yazdırır.

	// İki tamsayı eşit olana kadar döngüyü çalıştırır.
	for a != b {
		if b < a {
			b++ // b, a'dan küçükse, b'yi artırır.
		} else {
			b-- // b, a'dan büyükse, b'yi azaltır.
		}
		fmt.Print(" ", b) // Her adımda b'yi ekrana yazdırır ve bir boşluk ekler.
	}
	fmt.Println() // Döngü tamamlandığında yeni bir satıra geçer.
}
