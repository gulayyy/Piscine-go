package main

import "fmt"

func main() {
	// "donnie" adında bir "Pilot" türünden bir değişken tanımlanır.
	var donnie Pilot

	// "donnie" pilotun özellikleri atanır.
	donnie.Name = "Donnie"
	donnie.Life = 100.0
	donnie.Age = 24
	donnie.Aircraft = AIRCRAFT1

	// "donnie" pilotun bilgileri yazdırılır.
	fmt.Println(donnie)
}

// "AIRACRAFT1" sabit bir değere atanır.
const AIRCRAFT1 = 1

// "Pilot" adında bir yapı (struct) tanımlanır. Bu yapı, pilot bilgilerini temsil eder.
type Pilot struct {
	Name     string
	Life     int
	Age      int
	Aircraft int
}
