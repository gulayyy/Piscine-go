package piscine

func IterativePower(nb int, power int) int {
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
