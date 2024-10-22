package piscine

func IterativeFactorial(nb int) int {
	if nb == 0 || nb == 1 {
		return 1
	}
	if nb < 0 {
		return 0
	}
	if nb < 25 {
		faktoriyel := 1
		for i := 1; i <= nb; i++ {
			faktoriyel *= i
		}
		return faktoriyel
	} else {
		return 0
	}
}
