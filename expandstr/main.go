package main

import (
	"os"

	"github.com/01-edu/z01"
)

func Trim(str string) string {
	for i, r := range str {
		if r != ' ' {
			str = str[i:]
			break
		}
	}
	for i := len(str) - 1; i >= 0; i-- {
		if str[i] != ' ' {
			str = str[:i+1]
			break
		}
	}
	return str
}

func main() {
	if len(os.Args) == 2 {
		str := Trim(os.Args[1])
		result := ""
		count := 0
		for _, r := range str {
			if r != ' ' {
				result += string(r)
				count = 0
			} else {
				count++
				if count == 1 {
					result += "   "
				}
			}
		}
		for _, r := range result {
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
	} else {
		return
	}
}
