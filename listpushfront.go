package piscine

// sonradan gelen dökümanları başa alarak her seferinde başa gelen elemanı değiştiriyor
func ListPushFront(l *List, data interface{}) {
	// Veriyi içeren yeni bir NodeL elemanı oluşturulur.
	n := &NodeL{Data: data}

	// Eğer liste boşsa, yeni eleman liste başı (Head) olur ve işlem sonlandırılır.
	if l.Head == nil {
		l.Head = n
		return
	}

	// Eğer liste doluysa, yeni elemanın Next özelliği mevcut listenin başını işaret eder.
	n.Next = l.Head

	// Listenin başı (Head), yeni eklenen eleman olan 'n' olarak güncellenir.
	l.Head = n
}
