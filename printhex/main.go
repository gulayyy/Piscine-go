package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) !=2{// 1'den farklı argüman gelme durumunu kontrol ediyoruz
		return
	}
	deger := os.Args[1]

	sayi , err := strconv.Atoi(deger)//gelen degeri önce tam sayıya çeviriyoruz 
	if err != nil{
		fmt.Println("ERROR")
		return 
	}
	hex := fmt.Sprintf("%x",sayi)//burdaki x degeri 16'lık tabana sayı çevirmede kullanılır

	fmt.Println(hex)
}
