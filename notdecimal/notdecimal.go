package main

import (
	"fmt"
)

func main() {
	fmt.Print(NotDecimal("0.1"))
	fmt.Print(NotDecimal("174.2"))
	fmt.Print(NotDecimal("0.1255"))
	fmt.Print(NotDecimal("1.20525856"))
	fmt.Print(NotDecimal("-0.0f00d00"))
	fmt.Print(NotDecimal(""))
	fmt.Print(NotDecimal("-19.525856"))
	fmt.Print(NotDecimal("1952"))
}

func NotDecimal(dec string) string {
	bos := ""
	dizi := []rune(dec)
	for i, k := range dizi {
		if k >= 'a' && k <= 'z' || k >= 'A' && k <= 'Z' {
			bos = dec
			break
		}

		if k != '.' {
			if i == 0 && k == '0' {
				continue
			} else {
				bos += string(k)
			}
		}
	}
	return bos + "\n"
}
