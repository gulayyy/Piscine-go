package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	arg := os.Args[1]
	a, err := strconv.Atoi(arg)
	if err != nil {
		for i := 0; i < 8; i++ {
			z01.PrintRune('0')
		}
		return
	}

	var str string
	if a == 0 {
		str = "00000000"
	}

	for len(str) != 8 {
		str = string(rune(a%2+48)) + str
		a /= 2
	}

	for _, i := range str {
		z01.PrintRune(i)
	}
}
