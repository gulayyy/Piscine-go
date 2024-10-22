package main

import (
	"fmt"
)

func main() {
	a1 := []int{0, 1, 2, 3, 4, 5}
	a2 := []int{0, 2, 1, 3}
	a3 := []int{400968, 294978, -597489, 83667, -708071, -669730, -847367, 858658}

	result1 := IsSorted(f, a1)
	result2 := IsSorted(f, a2)
	result3 := IsSorted(f, a3)

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
}

func f(a, b int) int {
	if a > b {
		return 1
	} else if b > a {
		return -1
	}
	return 0
}
func IsSorted(f func(a, b int) int, a []int) bool {
	for i := 0; i < len(a)-1; i++ {
		if f(a[i], a[i+1]) == -1 {
			continue
		} else if f(a[i], a[i+1]) == 1 {
			return false
		}
	}
	return true
}
