package piscine

func RecursiveFactorial(nb int) int {
	if nb == 0 {
		return 1
	} else if nb < 0 || nb > 25 {
		return 0
	}
	return RecursiveFactorial(nb-1) * nb
}

/*if nb < 25 {
	if nb <= 1 {
		return 1
	}
	return nb * RecursiveFactorial(nb-1)
} else {
	return 0
}*/
