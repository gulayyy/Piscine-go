package main

import (
	"fmt"
)

func HashCode(dec string) string {
	bos := ""
	for _, i := range dec {
		if i < ' ' {
			bos += string((int(i+33) + len(dec)) % 127)
		} else {
			bos += string((int(i) + len(dec)) % 127)
		}
	}

	return bos
}

func main() {
	fmt.Println(HashCode("A"))
	fmt.Println(HashCode("AB"))
	fmt.Println(HashCode("BAC"))
	fmt.Println(HashCode("Hello World"))
}
