package piscine

func SplitWhiteSpaces(s string) []string {
	var kelimeler []string

	kelime := ""

	for _, a := range s {
		if a == ' ' || a == '\t' || a == '\n' {
			if kelime != "" {
				kelimeler = append(kelimeler, kelime)
				kelime = ""
			}
		} else {
			kelime += string(a)
		}
	}
	if kelime != "" {
		kelimeler = append(kelimeler, kelime)
	}
	return kelimeler
}
