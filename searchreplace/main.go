package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 4 {
		
	s1 := os.Args[1]
	s2 := os.Args[2]
	s3 := os.Args[3]
	output := ""

	for _, char := range s1 {
		if string(char) == s2 { // Dize içindeki her karakteri tek tek kontrol etmek için
			output += s3 // Eğer karakter s2 ile eşitse, onu s3 ile değiştir
		} else {
			output += string(char) // Aksi halde, karakteri değiştirmeden output'a ekle
		}
	}
	fmt.Println(output)
	}
}
