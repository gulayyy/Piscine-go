package piscine

func Swap(a *int, b *int) {
	*a, *b = *b, *a
}

//c :=*a yeniden değişken olarak tanımladığımız için arasına : koyuyoruz
//*a =*b
//*b :=c
