package piscine

func ListLast(l *List) interface{} {
	// Eğer liste boşsa, yani başlangıç elemanı (Head) yoksa, 'nil' değeri döndürülür.
	if l.Head == nil {
		return nil
	} else {
		// Eğer liste doluysa, son elemana ulaşmak için bir döngü kullanılır.
		// Başlangıç olarak, 'baslangic' adında bir işaretçi, listenin başını (Head) işaret eder.
		baslangic := l.Head

		// Döngü, 'current' elemanının Next özelliği 'nil' olana kadar devam eder.
		// Yani son elemana ulaşana kadar ilerlenir.
		for baslangic.Next != nil {
			baslangic = baslangic.Next
		}

		// Döngü sonunda, 'baslangic' son elemanı işaret eder ve bu elemanın verisi (Data) döndürülür.
		return baslangic.Data
	}
}
