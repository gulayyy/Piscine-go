package main

import (
	"strings"
	"unicode"
	"fmt"
)

func IsCapitalized(s string) bool {
	// Eğer string boş ise, false döndür
	if len(s) == 0 {
		return false
	}

	// Stringi kelimelere ayır
	words := strings.Fields(s)

	// Her kelime üzerinde döngü
	for _, word := range words {
		// Kelimenin ilk karakterinin büyük harf veya alfabedışı bir karakter olup olmadığını kontrol et
		firstChar := rune(word[0])
		if !unicode.IsUpper(firstChar) && !unicode.IsLetter(firstChar) {
			return false
		}
	}

	// Tüm kelimeler koşulları sağlıyorsa, true döndür
	return true
}

func main() {
	fmt.Println(IsCapitalized("Hello! How are you?"))
	fmt.Println(IsCapitalized("Hello How Are You"))
	fmt.Println(IsCapitalized("Whats 4this 100K?"))
	fmt.Println(IsCapitalized("Whatsthis4"))
	fmt.Println(IsCapitalized("!!!!Whatsthis4"))
	fmt.Println(IsCapitalized(""))
}