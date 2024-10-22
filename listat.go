package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	sayac := 0
	for l != nil { // içindeki l yi kontrol ediyor boş mu diye
		if sayac == pos { // bizden istenen index değerine üstünde dolaşırkenki index eşitse içine girip listedekini yazdırsın
			return l
		}
		l = l.Next // sonraki index değerine geçmek için kullanılır
		sayac++
	}
	return nil
}
