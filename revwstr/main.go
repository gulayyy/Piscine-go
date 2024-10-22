package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	deger := os.Args[1]

	kelimeler := strings.Fields(deger)

	var bos string
	for i := len(kelimeler) - 1; i >= 0; i-- {
		bos += kelimeler[i] + " "
	}
	fmt.Println(bos)

	son := ""
	for j := len(bos) - 1; j >= 0; j-- {
		if (j == len(bos)-1 && bos[j] == ' ') || (j == 0 && bos[j] == ' ') {
			continue
		} else {
			son = string(bos[j]) + son
		}
	}

	fmt.Println(son)

	// ya da izin veririse strings.Trimspace(bos) yaparak ba ve sondaki boşluşarı silebiliriz.
}
