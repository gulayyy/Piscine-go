package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	if len(os.Args) == 3 {
		arg1 := os.Args[1]
		arg2 := os.Args[2]

		bos := []rune{}

		for _,i:=range arg1{
			for _,k:=range arg2{
				if i == k && !strings.ContainsRune(string(bos),i){
					bos = append(bos, i)
				}
			}
		}
		fmt.Println(string(bos))
	}
}