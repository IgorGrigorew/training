package lessoncodebasic

// MostPopularWord(words []string) string,
// которая возвращает самое часто встречаемое слово в слайсе.
// Если таких слов несколько, то возвращается первое из них.
func MostPopularWord(words []string) string {
	mapa := make(map[string]int, 0)
	for _, vol := range words {
		_, ok := mapa[vol]
		if ok {
			mapa[vol] += 1
		} else {
			mapa[vol] = 1
		}
	}
	var result string
	var a int
	for key, volue := range mapa {
		if a < volue {
			a = volue
			result = key
		}
	}
	return result
}
