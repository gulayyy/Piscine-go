package main

import (
	"fmt"
	"os"
	"strconv"
)

// gcd, verilen iki tamsayının en büyük ortak bölenini (GCD) hesaplar ve döndürür.
func gcd(a, b int) int {
	for a != b {
		if a > b {
			a -= b // a, b'den büyükse, a'dan b'yi çıkarır.
		} else {
			b -= a // a, b'den küçükse, b'den a'yı çıkarır.
		}
	}
	return a // a ve b aynı olduğunda, en büyük ortak bölen (GCD) döndürülür.
}

func main() {
	if len(os.Args) == 3 {
		a, _ := strconv.Atoi(os.Args[1]) // İlk komut satırı argümanını tamsayıya dönüştürür.
		b, _ := strconv.Atoi(os.Args[2]) // İkinci komut satırı argümanını tamsayıya dönüştürür.
		fmt.Println(gcd(a, b))           // gcd fonksiyonunu kullanarak GCD'yi hesaplar ve ekrana yazdırır.
	}
}
