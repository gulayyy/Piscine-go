package piscine

func UltimateDivMod(a *int, b *int) {
	bolum := *a / *b
	kalan := *a % *b

	*a = bolum
	*b = kalan
}
