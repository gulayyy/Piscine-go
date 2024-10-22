package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for a := '0'; a <= '9'; a++ {
		for b := '0'; b <= '9'; b++ {
			d := b + 1                  // d ile b arasında bir fazlalık bir sayı farkı olduğu için d yi b ye bağlı bir değişken gibi düşündük
			for c := a; c <= '9'; c++ { // c ve a döngü içlerinde genelde aynı şekilde başladığı için ve projede tekrar etmesini istememişti tekrardan kaçm için bu şekilde bir başlangıç tercih ettik
				for ; d <= '9'; d++ {
					z01.PrintRune(a)
					z01.PrintRune(b)
					z01.PrintRune(' ')
					z01.PrintRune(c)
					z01.PrintRune(d)
					if a < '9' || b < '8' || c < '9' || d < '9' {
						z01.PrintRune(',') // En son kısımda bizden virgül konmasını istemediği için bu şekilde yaptık
						z01.PrintRune(' ') // En son kısımda bizden boşluk konmasını istemediği için bu şekilde yaptık
					}
				}
				d = '0' // d b ye bağlı bir değişken olarak tanımladığımız ve hep b den bir fazla olmasını istediğimiz için işlem her tekrarlandığında bu şekilde sıfırlama işlemi yaptık
			}
		}
	}
	z01.PrintRune('\n')
}
