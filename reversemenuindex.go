package piscine

// ReverseMenuIndex, verilen bir dilimdeki öğelerin sırasını tersine çevirir.
func ReverseMenuIndex(menu []string) []string {
	// Yeni bir boş dilim tanımlanır, bu dilim sonuçları içerecektir.
	var dizi []string

	// Dilimi sondan başa doğru tarayacak bir döngü başlatılır.
	// Başlangıç indeksi len(menu) - 1 olarak ayarlanır, çünkü indeksler sıfırdan başlar ve sondan başlayarak gitmek istiyoruz.
	for i := len(menu) - 1; i >= 0; i-- {
		// Her döngü adımında, menu dilimindeki öğenin tersine çevrilen sırası olan i indeksi kullanılarak bu öğe dizi dilimine eklenir.
		dizi = append(dizi, menu[i])
	}

	// Döngü tamamlandığında, dizi dilimi öğelerin sırasını tersine çevirmiş olarak doldurulur ve fonksiyon tarafından döndürülür.
	return dizi
}
