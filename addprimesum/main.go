package main

import (
	"fmt"
	"os"
	"strconv"
)

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func sumOfPrimes(input int) int {
	if input <= 0 {
		return 0
	}

	sum := 0
	for i := 2; i <= input; i++ {
		if isPrime(i) {
			sum += i
		}
	}
	return sum
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("0")
		return
	}

	num, err := strconv.Atoi(os.Args[1])
	if err != nil || num <= 0 {
		fmt.Println("0")
		return
	}

	fmt.Println(sumOfPrimes(num))
}