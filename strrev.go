package piscine

func StrRev(s string) string {
	uzunluk := StrLen(s)
	str := ""
	karakter := []rune(s)

	for i := uzunluk - 1; i >= 0; i-- {
		str += string(karakter[i])
	}
	return str
}
