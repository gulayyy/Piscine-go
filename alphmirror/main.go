package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 2 {
		arg := []rune(os.Args[1]) // Komut satırından alınan argümanı rune dilimine dönüştürür.

		// rune dilimini dolaşarak her harf için işlem yapar.
		for i, ch := range arg {
			if ch >= 'a' && ch <= 'z' {
				arg[i] = 'z' - ch + 'a' // Küçük harf ise, alfabe ters çevirme işlemi yapar.
			} else if ch >= 'A' && ch <= 'Z' {
				arg[i] = 'Z' - ch + 'A' // Büyük harf ise, alfabe ters çevirme işlemi yapar.
			}
		}

		fmt.Print(string(arg)) // Ters çevrilen rune dilimini ekrana yazdırır.
	}

	fmt.Println() // Bir satır atlayarak yeni satıra geçer.
}
