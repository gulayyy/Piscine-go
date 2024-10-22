package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	dizi := [10]int{}     // 10 elemanlı bir dizi oluşturduk
	if n < 10 && n >= 0 { // dizinin içi ful sıfır//tek basamak kontorlü yaaptık
		z01.PrintRune(rune(n + 48)) // n+48 ile int değeri karaktere çeviriyoz tam tersi durumda da karakteri int a çeviriyoz
	} else {
		for n > 0 {
			dizi[n%10]++ // dizinin n%10. karakterine git ve bir aartır
			n /= 10      // sonrasında verilen n değerini ile önce / işlemine sokuyoruz sonra % işlemine sokuyoruz ve o indekse gidip bir artırıyoruz//verilen değer sıfırdan küçük olana kadar işlem böyle devam ediyor
		}
		for index, value := range dizi {
			if value != 0 {
				for i := 0; i < value; i++ { // bu işlem sadece değer sıfır olduğunda dönüyor
					z01.PrintRune(rune(index + 48))
				}
			}
		}
	}
}
