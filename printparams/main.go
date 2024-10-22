package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguman := os.Args
	for i := 1; i < len(arguman); i++ {
		for _, a := range []rune(arguman[i]) {
			z01.PrintRune(a)
		}
		z01.PrintRune('\n')
	}
}
