package piscine

func ListReverse(l *List) { // verilen değerler arasında ikili ikili yer değiştirme yapar
	// İki geçici değişken tanımla: "a" (önceki düğüm) ve "current" (şu anki düğüm)
	var a, current, next *NodeL
	current = l.Head // "current"ı bağlı listenin başlangıç düğümüne ayarla
	l.Tail = l.Head  // "l.Tail"i başlangıç düğümüne ayarla (bağlı listenin sonu da artık başa eşit olacak) yerlerini değiştirmiş olacak böylelikle

	// Bir döngü kullanarak bağlı listeyi tersine çevir
	for current != nil {
		next = current.Next // Bir sonraki düğümü "next" değişkenine kopyala
		current.Next = a    // "current" düğümünün "Next" alanını önceki düğüme (a) ayarla
		a = current         // "a" değişkenini şu anki düğümün üzerine taşı
		current = next      // "current" değişkenini bir sonraki düğüme taşı
	}

	// Bağlı listenin başını, son düğüme yani "a" değişkenine ayarla
	l.Head = a
}
