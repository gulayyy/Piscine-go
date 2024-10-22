package piscine

func CompStr(a, b interface{}) bool {
	// Bu fonksiyon, iki arayüz türünden argümanı karşılaştırır.
	return a == b
}

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	// Bağlı listeyi baştan sona dolaşmak için iterator (iterator) oluşturulur.
	baslangic := l.Head

	// Döngü, iterator'ün null olmadığı sürece çalışır (bağlı listenin sonuna kadar dolaşılır).
	for baslangic != nil {
		// Mevcut düğümün verisi (baslangic.Data) ve referans (ref) comp işlevi kullanılarak karşılaştırılır.
		if comp(baslangic.Data, ref) {
			// Eğer karşılaştırma sonucunda eşleşme bulunursa, bu düğümün verisine işaret eden bir işaretçi döndürülür.
			return &baslangic.Data
		}
		// Bir sonraki düğüme geçmek için iterator ilerletilir.
		baslangic = baslangic.Next
	}
	// Hiçbir eşleşme bulunmazsa veya bağlı liste boşsa, fonksiyon null döndürür.
	return nil
}

// üzerinde dolaşıyor ilk adresini tanımlıyor sonrasında da kendini yazdırıyor
