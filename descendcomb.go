package piscine

import (
	"github.com/01-edu/z01"
)

func DescendComb() {
	for k := '9'; k >= '0'; k-- {
		for m := '9'; m >= '0'; m-- {
			for i := '9'; i >= '0'; i-- {
				for j := '9'; j >= '0'; j-- {
					if k > i || (k == i && m > j) {
						z01.PrintRune(k)
						z01.PrintRune(m)
						z01.PrintRune(' ')
						z01.PrintRune(i)
						z01.PrintRune(j)

						if k == '0' && m == '1' && i == '0' && j == '0' {
						} else {
							z01.PrintRune(',')
							z01.PrintRune(' ')
						}
					}
				}
			}
		}
	}
}
