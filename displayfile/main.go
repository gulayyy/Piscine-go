package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// Hatayı kontrol etmek için fonksiyonumuz
func kontrol(hata error) {
	if hata != nil {
		panic(hata) // hata printrune gibi bir şey hata gördüğü zaman direkmen programı durduruyor
	}
}

func main() {
	arg := os.Args[1:]
	uzunluk := len(arg)

	if uzunluk == 0 {
		fmt.Println("File name missing")
	} else if uzunluk > 1 {
		fmt.Println("Too many arguments")
	} else {
		// Okunacak dosyamızı belirtiyoruz
		dosya, hata := ioutil.ReadFile("quest8.txt")
		// Hata kontrolü yapıyoruz.
		kontrol(hata)
		// Dosyamızın içeriğini ekrana bastırıyoruz.
		fmt.Print(string(dosya))
	}
}
