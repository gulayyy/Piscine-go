package piscine

func CollatzCountdown(start int) int { // Collatz  varsayımını kullandık
	if start <= 0 { // Burda sıfırdan küçüklük ve başlangıç sıfıra eşit olduğundaki kısmı kontrol ediyoruz
		return -1
	}
	adim := 0
	for start != 1 { // Birden başladığı zaman ilk önce teklik çiftlik kontrolü yaptık
		if start%2 == 0 {
			start /= 2 // eğer çift ise yarısını alıyoruz
		} else {
			start = start*3 + 1 // eğer tekse üç katının bir fazlasını alıyoruz
		}
		adim++ // adımını da tek tek artıyoruz
	}
	return adim
} // Ve bu döngü en son 1 sayısını bulana kadar devam ediyor
