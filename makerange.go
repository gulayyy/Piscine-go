package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil // değişkenin herhangi bir değer almadığı durumlar için kullanılır
	}
	boyut := max - min
	sayilar := make([]int, boyut) // makein appendden farkı gelen boyut kadar uzunlukta açması
	for i := 0; i < boyut; i++ {
		sayilar[i] = min + i // dizinin nerden başladığını buluyoruz sonrasında farkı buluyoruz fark kAdar döndürüp başlangıç değerine ekliyoruz ve max kısmına eşit olana kadar döndürüyoruz
	}
	return sayilar
}
