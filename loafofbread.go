package piscine

// LoafOfBread, verilen bir dizeyi belirli bir desenle biçimlendiren bir fonksiyondur.
func LoafOfBread(str string) string {
	// Eğer gelen dize boşsa, yeni bir satır karakteri döndürülür.
	if str == "" {
		return "\n"
	}

	// Eğer gelen dize 5 karakterden kısa ise, "Invalid Output" mesajı döndürülür.
	if len(str) < 5 {
		return "Invalid Output\n"
	}

	// Sonuç dizesi için bir byte dilimi oluşturulur. Bu dilim, gelen dizeyi biçimlemek için kullanılacaktır.
	result := make([]byte, 0, len(str)+len(str)/5)

	// Bir sayac (j) başlatılır.
	j := 0

	// Gelen dizeyi karakter karakter dolaşmak için bir döngü başlatılır.
	for i := 0; i < len(str); i++ {
		// Eğer hala 5 karakteri doldurmadıysak ve mevcut karakter bir boşluksa, bu karakteri atlar ve bir sonraki karaktere geçeriz.
		if j < 5 && rune(str[i]) == ' ' {
			continue
		}

		// Eğer sayac (j) 5'e ulaşırsa, bir boşluk karakteri eklenir.
		if j == 5 {
			// Eğer şu anki karakter boşluk ise ve bir sonraki karakter de boşluksa, bu iki boşluğu atlar ve bir sonraki karaktere geçeriz.
			if i != len(str)-1 && str[i+1] == ' ' {
				continue
			}

			// Eğer şu anki karakter son karakter ise, döngüden çıkılır.
			if i == len(str)-1 {
				break
			}

			// Sonuç dizesine bir boşluk eklenir ve sayac (j) sıfırlanır.
			result = append(result, ' ')
			j = 0
			continue
		}

		// Eğer yukarıdaki koşullar sağlanmazsa, bu karakter sonuç dizesine eklenir ve sayac (j) artırılır.
		result = append(result, str[i])
		j++
	}

	// Sonuç dizesine bir yeni satır karakteri eklenir ve sonuç dize döndürülür.
	result = append(result, '\n')
	return string(result)
}
