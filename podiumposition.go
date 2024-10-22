package piscine

// PodiumPosition, bir yarışın podyum sıralamasını düzenler ve döndürür.
func PodiumPosition(podium [][]string) [][]string {
	// Ödül pozisyonlarını (1., 2. ve 3. gibi) indekslere karşılık gelecek şekilde bir harita (map) oluşturulur.
	positions := map[string]int{
		"1st Place": 0,
		"2nd Place": 1,
		"3rd Place": 2,
		"4th Place": 3,
	}

	// İç içe iki döngü kullanılarak podyumun sıralanması yapılır.
	for i := 0; i < len(podium); i++ {
		for j := i + 1; j < len(podium); j++ {
			// İç içe döngüler, yarışçıların sıralarını karşılaştırır ve yer değiştirmeleri gerekiyorsa yer değiştirir.
			if positions[podium[i][0]] > positions[podium[j][0]] {
				podium[i], podium[j] = podium[j], podium[i]
			}
		}
	}

	// Sıralanmış podyumu döndürür.
	return podium
}
