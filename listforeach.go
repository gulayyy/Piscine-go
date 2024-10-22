package piscine

// ListForEach, verilen bağlı listeyi baştan sona dolaşarak her düğüm üzerine
// bir işlem uygulayan bir fonksiyondur. Her döngü adımında, verilen işlev (f)
// mevcut düğüme (baslangic) uygulanır.
func ListForEach(l *List, f func(*NodeL)) {
	baslangic := l.Head // Bağlı listenin başından başla
	for baslangic != nil {
		f(baslangic)               // Belirtilen işlemi mevcut düğüme uygula
		baslangic = baslangic.Next // Bir sonraki düğüme geç
	}
}

// Add2_node, verilen bir NodeL düğümünün verisine bağlı olarak işlem yapar.
// Eğer veri bir tamsayı (int) ise, bu tamsayıya 2 ekler. Eğer veri bir metin (string)
// ise, bu metnin sonuna "2" ekler.
func Add2_node(node *NodeL) {
	switch node.Data.(type) {
	case int:
		node.Data = node.Data.(int) + 2
	case string:
		node.Data = node.Data.(string) + "2"
	}
}

// Subtract3_node, verilen bir NodeL düğümünün verisine bağlı olarak işlem yapar.
// Eğer veri bir tamsayı (int) ise, bu tamsayıdan 3 çıkarır. Eğer veri bir metin (string)
// ise, bu metnin sonuna "-3" ekler.
func Subtract3_node(node *NodeL) {
	switch node.Data.(type) {
	case int:
		node.Data = node.Data.(int) - 3
	case string:
		node.Data = node.Data.(string) + "-3"
	}
}
