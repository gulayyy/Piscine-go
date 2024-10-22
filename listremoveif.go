package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	temp := l.Head
	prev := l.Head

	// İlk döngü: Başındaki düğümleri verilen referans ile eşleşenlerin hepsini kaldır.
	for temp != nil && temp.Data == data_ref {
		l.Head = temp.Next // Başlangıç düğümünü kaldır ve başı güncelle
		temp = l.Head      // Geçici işaretçiyi başa geri al
	}

	// İkinci döngü: Geriye kalan düğümleri kontrol ederek verilen referansa sahip olanları kaldır.
	for temp != nil {
		for temp != nil && temp.Data != data_ref {
			prev = temp      // Önceki düğümü güncelle
			temp = temp.Next // Geçici işaretçiyi bir sonraki düğüme ilerlet
		}

		// Eğer temp artık boşsa, yani bağlı listenin sonu ise, işlemi sonlandır.
		if temp == nil {
			return
		}

		// Eşleşme bulunduğunda:
		// - Önceki düğümün Next'i, eşleşen düğümün Next'ine bağlanır (eşleşen düğüm atlanmış olur).
		// - Geçici işaretçi bir sonraki düğüme ilerletilir.
		prev.Next = temp.Next
		temp = temp.Next
	}
}
