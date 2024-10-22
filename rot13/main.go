package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	arg := os.Args[1]
	for _, i := range arg {
		i = rot13(i)
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')
}

func rot13(x rune) rune {
	if x >= 'a' && x <= 'm' || x >= 'A' && x <= 'M' {
		return x + 13
	} else if x >= 'n' && x <= 'z' || x >= 'N' && x <= 'Z' {
		return x - 13
	} else if x == ' ' {
		return x
	}
	return x
}
