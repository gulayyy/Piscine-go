package main

import (
	"fmt"
	"os"
)
func hiddenp(a , b string)bool{
	if a == ""{
		return false
	}
	i := 0
	for _, char := range b{
		if byte(char) == a[i]{
			i++
			if i == len(a){
				return true
			}
		}
	}
	return false
}
func main(){
	if len(os.Args)== 3{
		s1 := os.Args[1]
		s2 := os.Args[2]
		if hiddenp(s1,s2){
			fmt.Println("1")
		}else{
			fmt.Println("0")
		}
	}
}
