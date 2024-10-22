package piscine

func AppendRange(min, max int) []int {
	sayilar := []int{}

	if min < max {
		for i := min; i < max; i++ {
			sayilar = append(sayilar, i)
		}
		return sayilar
	}
	return nil // değişkenin herhangi bir değer almadığı durumlar için kullanılır
}
