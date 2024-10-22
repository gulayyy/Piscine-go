package piscine

type NodeL struct { // Her bir bağlı liste elemanının verisini ve bir sonraki elemanı işaret eden Next'i içeren NodeL yapısı tanımlanır.
	Data interface{} // NodeL yapısının bir alanını (field) tanımlar ve bu alana çeşitli türlerde veri (data) depolayabilme yeteneği verir
	Next *NodeL
}

type List struct {
	Head *NodeL // liste başını ifade eder
	Tail *NodeL // liste sonunu ifade eder
}

func ListPushBack(l *List, data interface{}) {
	// Verilen veriyi içeren yeni bir NodeL oluşturuluyor.
	Node := &NodeL{Data: data}

	// Eğer liste boşsa, yeni Node, liste başı (Head) olur.
	if l.Head == nil {
		l.Head = Node
	} else {
		// Eğer liste doluysa, mevcut listedeki son elemana ulaşmak için bir döngü kullanılır.
		// Bu döngü, listenin sonuna gitmek için kullanılır.
		deger := l.Head
		for deger.Next != nil {
			deger = deger.Next
		}

		// Döngü sonunda, current, listenin sonundaki Node'u işaret eder.
		// Yeni Node, bu son Node'un Next özelliğine eklenir, böylece yeni Node
		// liste sonu (Tail) haline gelir.
		deger.Next = Node
	}
}
