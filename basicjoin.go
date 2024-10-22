package piscine

func BasicJoin(elems []string) string {
	bos := ""
	for _, i := range elems {
		bos += i
	}
	return bos
}
