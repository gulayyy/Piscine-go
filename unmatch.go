package piscine

func Unmatch(a []int) int {
	for _, i := range a { // Dilimdeki her bir elemanın üzerinde dolaşır
		sayac := 0
		for _, b := range a { // Dilimde ele aldığımız eleman haricci diğer elemanların üzerinde dolaşmka için kullanılır
			if i == b { // eğer eşit sayı bulursa sayacı bir artırarak içinde eş sayı buldum diye artırır
				sayac++
			}
		}
		if sayac == 1 || sayac%2 == 1 { // eğer sayacın içinde her bir eleman eşleşmiyorsa eşleşen sayıyı döndürür mesela bir sayıdan 3 tane ise kalan 1 taneyi dödürür çünkü şey der bu sayının eşi yok deyip yollar
			return i
		}
	}
	return -1 // eğer dilim içiçndeki bütün sayıların bir eşi varsa -1 döndürür
}
