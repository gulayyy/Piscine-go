package piscine

func ListForEachIf(l *List, f func(*NodeL), cond func(*NodeL) bool) {
	// Bağlı listenin başından başlamak için bir iterator (it) oluştur.
	it := l.Head

	// Düğümleri dolaşmak için bir döngü başlatılır.
	for it != nil {
		// "cond" işlevi, belirli bir koşulu kontrol eder. Eğer bu koşul sağlanıyorsa, yani "cond(it)" true dönerse:
		if cond(it) {
			// "f" işlevi, mevcut düğüme (it) uygulanır. Yani, bu işlev mevcut düğüm üzerinde belirli bir işlemi gerçekleştirir.
			f(it)
		}
		// Bir sonraki düğüme geçmek için iterator ilerletilir.
		it = it.Next
	}
}

func IsPositiveNode(node *NodeL) bool {
	// Fonksiyon, bir NodeL düğümünün veri türünü kontrol eder.
	switch node.Data.(type) {
	case int, float32, float64, byte:
		// Veri alanı bir tamsayı, kayan noktalı sayı veya bayt türündeyse:
		// Veri alanındaki sayının pozitif olup olmadığını kontrol eder.
		return node.Data.(int) > 0
	default:
		// Eğer veri alanı yukarıdaki türlerden birine sahip değilse (örneğin, bir dize):
		// Daima false döner, çünkü sadece belirli türler için pozitiflik kontrolü yapar.
		return false
	}
}

func IsAlNode(node *NodeL) bool {
	// Fonksiyon, bir NodeL düğümünün veri türünü kontrol eder.
	switch node.Data.(type) {
	case int, float32, float64, byte:
		// Veri alanı bir tamsayı, kayan noktalı sayı veya bayt türündeyse:
		// Bu türlerin düğümleri "false" döner, çünkü bunlar işaretsiz türlerdir.
		return false
	default:
		// Eğer veri alanı yukarıdaki türlerden birine sahip değilse (örneğin, bir dize):
		// Daima true döner, çünkü sadece belirli türlerin "false" dönmesini sağlar.
		return true
	}
}
