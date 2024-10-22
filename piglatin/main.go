package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Tek bir argümanın sağlandığını kontrol ediyoruz
	if len(os.Args) != 2 {
		return
	}

	// Giriş kelimesini komut satırı argümanından alıyoruz
	kelime := os.Args[1]

	// Karşılaştırma yaparken büyük-küçük harf duyarlılığını ortadan kaldırmak için kelimeyi küçük harfe çeviriyoruz
	kelime = strings.ToLower(kelime)

	// Kelimenin içinde sesli harf var mı kontrol ediyoruz
	sesliVar := false
	for _, char := range kelime {
		if sesliMi(char) {
			sesliVar = true
			break
		}
	}

	// Eğer kelimenin içinde sesli harf yoksa, "Sesli harf yok" yazdırıyoruz
	if !sesliVar {
		fmt.Println("Sesli harf yok")
		return
	}

	// Kelimenin bir sesli harfle başlayıp başlamadığını kontrol ediyoruz
	if sesliMi(rune(kelime[0])) {
		// Eğer kelime bir sesli harfle başlıyorsa, sadece sonuna "ay" ekliyoruz
		fmt.Println(kelime + "ay")
		return
	}

	// Eğer kelime bir sessiz harfle başlıyorsa, ilk sesli harfe kadar olan bölümü sona taşıyıp "ay" ekliyoruz
	var onEk string
	var sonEk string
	for _, char := range kelime {
		if !sesliMi(char) {
			onEk += string(char)
		} else {
			sonEk = kelime[len(onEk):]
			break
		}
	}
	fmt.Println(sonEk + onEk + "ay")
}

// Bir karakterin sesli harf olup olmadığını kontrol eden fonksiyon
func sesliMi(char rune) bool {
	return char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u'
}
