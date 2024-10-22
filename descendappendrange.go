package piscine

func DescendAppendRange(max, min int) []int {
	sayilar := []int{}

	if min < max {
		for i := max; i > min; i-- {
			sayilar = append(sayilar, i)
		}
	}
	return sayilar
}
