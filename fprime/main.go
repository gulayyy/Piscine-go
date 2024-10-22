package main

import (
	"fmt"
	"os"
	"strconv"
)

// fprime, bir tamsayının asal çarpanlarını ekrana yazdırır.
func fprime(value int) {
	if value == 1 {
		return // Eğer girdi 1 ise, herhangi bir işlem yapmadan geri döner.
	}
	bolensayi := 2 // Başlangıçta bölen olarak 2'yi kullanırız.
	for value > 1 {
		if value%bolensayi == 0 { // Eğer value, bolensayi'a bölünebilirse,
			fmt.Print(bolensayi)      // Böleni ekrana yazdırırız.
			value = value / bolensayi // value'yi bölenin sonucuyla güncelleriz.

			if value > 1 {
				fmt.Print("*") // Eğer value hala 1'den büyükse, çarpma işaretini yazdırırız.
			}
			bolensayi-- // Bölme işlemini tekrar kontrol etmek için böleni azaltırız.
		}
		bolensayi++ // Bir sonraki böleni kontrol etmek için böleni artırırız.
	}
	fmt.Println() // Çarpanları yazdırdıktan sonra yeni bir satıra geçeriz.
}

func main() {
	if len(os.Args) == 2 { // Komut satırından sadece bir argüman alındığını kontrol eder.
		if i, err := strconv.Atoi(os.Args[1]); err == nil { // Argümanı bir tamsayıya dönüştürür.
			fprime(i) // fprime fonksiyonunu çağırarak çarpanları hesaplar ve ekrana yazdırır.
		}
	}
}
