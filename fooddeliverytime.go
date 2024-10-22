package piscine

type food struct {
	preptime int
}

func FoodDeliveryTime(order string) int {
	if order == "burger" {
		return 15
	} else if order == "chips" {
		return 10
	} else if order == "nuggets" {
		return 12
	}
	return 404
}

/*
type food struct {
	preptime int
	//price ile fiyat adlı bir değişken oluşturduğumuz zaman onu da aşağı kısımda kullanabiliriz
}
func FoodDeliveryTime(order string) int {
	var burger food//struct yapısı ile kendi eğişken yapımızı oluşturup food içerisinden çekiyourz
	burger.preptime = 15 // burger.price = 10
	var chips food
	chips.preptime = 10
	var nuggets food
	nuggets.preptime = 12
	if order == "burger" {
		return burger.preptime
	}

	if order == "chips" {
		return chips.preptime
	}
	if order == "nuggets" {
		return nuggets.preptime
	}
	return 404
}*/
