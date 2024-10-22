package piscine

func ListSize(l *List) int { // Gelen değerlerin toplaımını döndürüyor
	sayac := 0
	current := l.Head
	for current != nil {
		sayac++
		current = current.Next
	}
	return sayac
}
