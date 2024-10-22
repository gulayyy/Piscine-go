package piscine

func Atoi(s string) int {
	c := 0
	a := 1
	for index, i := range s {
		if index == 0 && i == '-' {
			a = -1
		} else if index == 0 && i == '+' {
			a = 1
		} else if i < '0' || i > '9' {
			return 0
		} else {
			c = c*10 + int(i-'0') // i burda ascıı kodundan çekiyor bir önceki sayıyı çekiyor bu formülde de c 10 ile çarpılaraak basamak basamak artırıyor
		}
	}
	return c * a
}
