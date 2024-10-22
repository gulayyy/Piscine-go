package main

import (
	"fmt"
)
func main() {
	fmt.Print(ReverseSecondHalf("This is the 1st half This is the 2nd half"))
	fmt.Print(ReverseSecondHalf(""))
	fmt.Print(ReverseSecondHalf("Hello World"))
}

func ReverseSecondHalf(str string) string {
	reverse := ""
	if str == "" {
		return "Invalid Output" + "\n"
	}
	if len(str) == 1 {
		return str
	}

	if len(str)%2 == 0 {
		reverse = str[(len(str) / 2):]
	} else {
		reverse = str[(len(str)+1)/2-1:]
	}
	bos := ""
	for i := len(reverse) - 1; i >= 0; i-- {
		bos += string(reverse[i])
	}
	return bos + "\n"
}