package piscine

func StringToIntSlice(str string) []int {
	dizi := []rune(str)
	var bos []int
	for _, i := range dizi {
		bos = append(bos, int(i))
	}
	return bos
}
