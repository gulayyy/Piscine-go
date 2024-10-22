package piscine

func Sqrt(nb int) int {
	if nb <= 0 {
		return 0
	}
	sqrt := 1
	for sqrt*sqrt < nb {
		sqrt++
	}
	if sqrt*sqrt == nb {
		return sqrt
	}
	return 0
}

/*func IterativePoweree(nb int, power int) int {
	if power < 0 {
		return 0
	}
	if power == 0 {
		return 1
	} else {
		sonuc := 1 // döngünün içinde veya dışında tanımlanması bir şeyi değiştirmiyor
		for i := 0; i < power; i++ {
			sonuc = sonuc * nb
		}
		return sonuc
	}
}
func Sqrt(nb int) int {
	a := 1
	for nb != IterativePoweree(a,2){
		a++
		if a> nb {
			return 0
		}
	}
	if nb == IterativePoweree(a,2){
		return a
	}
	return 0
}*/
