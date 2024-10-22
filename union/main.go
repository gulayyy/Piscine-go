package main

import (
	"fmt"
	"os"
	"strings"
)
func main(){
	if len(os.Args) == 3{
		var bos string
		deger1 := os.Args[1]
		deger2 := os.Args[2]

		for _,i:= range deger1{
			if !strings.ContainsRune(bos, i){//ilk gelen degeri daha önce bos içinde bulundu mu diye bakıyor eger bulunmamıssa ekliyor
				bos += string(i)
			}
		}
		for _,j := range deger2{
			if !strings.ContainsRune(bos,j){
				bos += string(j)
			}
			
		}
		fmt.Print(bos)
	}
	fmt.Println()
}
