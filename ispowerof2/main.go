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
	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		return
	}
	if num < 0 {
		return
	}
	if ispowerof2(num) {
		for _, r := range "true\n" {
			z01.PrintRune(r)
		}
	} else {
		for _, r := range "false\n" {
			z01.PrintRune(r)
		}
	}
}

func ispowerof2(n int) bool {
	if n == 1 {
		return true
	}
	if n < 1 {
		return false
	}
	cift := 1
	for cift <= n {
		if cift == n {
			return true
		}
		cift = cift * 2
	}
	return false
}
