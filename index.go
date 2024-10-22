package piscine

func Index(s string, toFind string) int {
	a := len(s)
	b := len(toFind)
	if a < b {
		return -1
	}
	for i := 0; i < a-b; i++ { // a-b ile son kısmında kalan harf sayısına göre arayıp aramaması gerektiğine karar veriyor
		if s[i:i+b] == toFind { // indexin bir özelliğidir verdiği ve bulmamızı istediğiparçanın sayısı kadar atlamaya çalışıyoruz
			return i
		}
	}
	return -1
}
