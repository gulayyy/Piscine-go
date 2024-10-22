package piscine

func JumpOver(str string) string {
	dizi := []rune(str)
	a := ""
	if len(dizi) > 2 {
		for i := 2; i < len(dizi); i++ {
			if i%3 == 2 {
				a += string(dizi[i])
			}
		}
		a += "\n"
	} else {
		a = "\n"
	}

	return a
}
