package main

import "fmt"

// FoldInt, verilen bir işlevi bir tamsayı dilimi üzerinde uygular, sonucu kaydeder ve yazdırır.
func FoldInt(f func(int, int) int, slice []int, acc int) {
	for _, num := range slice {
		acc = f(acc, num)
		fmt.Println(acc)
	}
}

// Toplama işlemi için fonksiyon.
func Add(a, b int) int {
	return a + b
}

// Çıkarma işlemi için fonksiyon.
func Sub(a, b int) int {
	return a - b
}

// Çarpma işlemi için fonksiyon.
func Mul(a, b int) int {
	return a * b
}

func main() {
	table := []int{1, 2, 3}
	ac := 93
	FoldInt(Add, table, ac)
	FoldInt(Mul, table, ac)
	FoldInt(Sub, table, ac)
	fmt.Println()

	table = []int{0}
	FoldInt(Add, table, ac)
	FoldInt(Mul, table, ac)
	FoldInt(Sub, table, ac)
}