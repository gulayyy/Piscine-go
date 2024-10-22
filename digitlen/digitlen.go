package main

import "fmt"

func DigitLen(n, base int) int {
	// Base 2 ile 36 arasında olmalıdır
	if base < 2 || base > 36 {
		return -1
	}

	// Eğer n negatifse, işareti tersine çevir ve rakamları say
	if n < 0 {
		n = -n
	}

	// Bölme işlemiyle rakamları say
	var count int
	for n > 0 {
		n /= base
		count++
	}
	return count
}
func main() {
	fmt.Println(DigitLen(100, 10))
	fmt.Println(DigitLen(100, 2))
	fmt.Println(DigitLen(-100, 16))
	fmt.Println(DigitLen(100, -1))
}