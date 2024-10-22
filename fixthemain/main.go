package main

import "github.com/01-edu/z01"

type Door struct {
	state bool
}

// PrintStr, verilen bir dizeyi rune'lere dönüştürerek ekrana yazdıran bir yardımcı fonksiyondur.
func PrintStr(str string) {
	arrayRune := []rune(str)
	for _, s := range arrayRune {
		z01.PrintRune(s)
	}
	z01.PrintRune(10) // 10, yeni satır karakterinin rune değerini temsil eder.
}

// OpenDoor, kapıyı açan bir fonksiyondur. Kapı durumu "true" olarak ayarlanır.
func OpenDoor(ptrDoor *Door) bool {
	PrintStr("Door Opening...")
	ptrDoor.state = true
	return true
}

// CloseDoor, kapıyı kapatan bir fonksiyondur. Kapı durumu "false" olarak ayarlanır.
func CloseDoor(ptrDoor *Door) bool {
	PrintStr("Door Closing...")
	ptrDoor.state = false
	return true
}

// IsDoorOpen, kapının açık olup olmadığını kontrol eden bir fonksiyondur.
func IsDoorOpen(ptrDoor *Door) bool {
	PrintStr("is the Door opened ?")
	return ptrDoor.state == true
}

// IsDoorClose, kapının kapalı olup olmadığını kontrol eden bir fonksiyondur.
func IsDoorClose(ptrDoor *Door) bool {
	PrintStr("is the Door closed ?")
	return ptrDoor.state == false
}

func main() {
	door := &Door{} // Door türünden bir nesne oluşturulur.

	OpenDoor(door) // Kapı açılır.

	if IsDoorClose(door) { // Eğer kapı kapalıysa, tekrar açılır
		OpenDoor(door)
	}

	if IsDoorOpen(door) { // Eğer kapı açıksa, kapatılır.
		CloseDoor(door)
	}

	if door.state == true { // Kapının durumu kontrol edilir ve açıksa kapatılır.
		CloseDoor(door)
	}
}
