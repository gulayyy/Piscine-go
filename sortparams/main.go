package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	AsciiSort(args)
	for _, arg := range args {
		PrintName(arg)
		z01.PrintRune('\n')
	}
}

func AsciiSort(arr []string) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if IsGreater(arr[j], arr[j+1]) {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func IsGreater(a, b string) bool {
	for i := 0; i < len(a) && i < len(b); i++ {
		if a[i] > b[i] {
			return true
		} else if a[i] < b[i] {
			return false
		}
	}
	return len(a) > len(b)
}

func PrintName(arg string) {
	for _, ch := range arg {
		z01.PrintRune(ch)
	}
}
