package piscine

func Compact(ptr *[]string) int {
	// İlk olarak, işlev bir dilim adresini alır ve bu dilimi "slice" adlı bir değişkene atar.
	slice := *ptr

	// Dilimin uzunluğunu hesaplar.
	length := len(slice)

	// Sıkıştırılmış (zero değerlere sahip öğelerin kaldırıldığı) bir dilim oluşturur.
	// Başlangıçta bu sıkıştırılmış dilim boş olacaktır.
	compactedSlice := make([]string, 0)

	// Döngü ile orijinal dilimdeki öğeleri tarar.
	for i := 0; i < length; i++ {
		// Eğer öğe boş bir dize değilse (sıfır olmayan bir değerse),
		// o öğeyi sıkıştırılmış dilime ekler.
		if slice[i] != "" {
			compactedSlice = append(compactedSlice, slice[i])
		}
	}

	// Sıkıştırılmış dilimi orijinal dilimin yerine atanır.
	*ptr = compactedSlice

	// Son olarak, sıkıştırılmış dilimin uzunluğu (sıfır olmayan öğelerin sayısı) döndürülür.
	return len(compactedSlice)
}

//
