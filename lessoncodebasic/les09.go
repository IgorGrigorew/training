package lessoncodebasic

import (
	"strings"
	"unicode"
)

//  возвращает только латинские символы из строки s
func LatinLetters(s string) string {
	bs := &strings.Builder{}
	for _, i := range s {
		if unicode.Is(unicode.Latin, i) {
			// записываем в bs руну
			bs.WriteRune(i)
		}
	}
	return bs.String()
}

func LatinLetters1(s string) string {
	var res string
	for _, vol := range s {
		if unicode.Is(unicode.Latin, vol) {
			res += string(vol)
		}
	}
	return res
}
