package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		return
	}

	str1 := os.Args[1]
	str2 := os.Args[2]

	if wdmatch(str1, str2) {
		fmt.Println(str1)
	}
}
func wdmatch(a, b string) bool {
	i := 0
	for _, k := range b {
        if i <len(a) && k == rune(a[i]){
            i++
        }
	}
    return i == len(a)
}
