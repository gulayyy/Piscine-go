package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	count := ""               // her  bir malzeme adını atıyor
	m := make(map[string]int) // m adında bir map oluşturulur. Bu map, her bir ürünün adını ve bu ürünün kaç kez görüldüğünü saklayacaktır
	for _, ch := range str {  // alışveriş fişi dizesini (str) karakter karakter tarıyor
		if ch == ' ' { // Her karakter kontrol edilir. Eğer karakter bir boşluk karakteri (' ') ise, if ch == ' ' koşulu sağlanır. Bu, bir ürün adının sona erdiğini ve adet bilgisinin başlayacağını gösterir.
			if count != " " { // Eğer count adında geçici dize boşluk karakteri değilse, bu, bir ürün adı olarak kabul edilir ve m haritasında bu ürün adının sayacı artırılır.
				m[count]++
				count = ""
			}
		} else { // Eğer karakter boşluk karakteri değilse, o karakter count dizesine eklenir. Bu, bir ürün adının karakterlerini bir araya getirir.
			count = count + string(ch)
		}
	}
	if count != " " {
		m[count]++
	}
	return m
}
