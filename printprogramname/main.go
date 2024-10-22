package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguman := os.Args[0] // os içinddeki argümanları çekmemize yardımcı oluyor
	dizi := []rune(arguman)

	for _, i := range dizi[2:] { // dosya adındaki almasını istemediğimiz kısımlardan sonraki kısımdan başlattık
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')
}
