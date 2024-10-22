package piscine

func Map(f func(int) bool, a []int) []bool {
	r := []bool{}
	for _, s := range a {
		r = append(r, f(s)) // her bir girilen değer için çıktı farklı olduğundan dolayı her bir çıktıyı alıp boş bir listeye atıyor ve her birini yazdırıyor
	}
	return r
}
