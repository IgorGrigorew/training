package lessonyandex

import (
	
	"slices"
	"sort"
	"strings"
	"unicode"
)



func StringLength(input string) int {
	return len(input)
}

//_____________________________________

func ConcatenateStrings(str1, str2 string) string {
	return str1 + " " + str2
}

//_______________________________________

func isLatin(input string) bool {

	for _, s := range input {
		if s > unicode.MaxASCII {
			return false
		}
	}

	return true
}

//______________________________________

func IsPalindrome(input string) bool {
	str := strings.ReplaceAll(input, " ", "")

	for i, y := 0, len(str)-1; i < y; i++ {

		if str[i] != str[y] {
			return false
		}

		y--
	}
	return true
}

//________________________________________

func Permutations(input string) []string {

		m := make(map[string]struct{})
	m[input] = struct{}{}
l := len(input)
by1 := []byte(input)
for i := 0; i < l; i++ {
	
	
	for y := l-1; y > 0; y-- {
		by := slices.Clone(by1)
		
		by[1], by[y] = by[y], by[1]
		_, key := m[string(by)]
		if !key {
			m[string(by)] = struct{}{}
		}

	}


by1[0], by1[l-1-i] = by1[l-1-i], by1[0]

}


res := make([]string, 0, len(m))
	for i , _ := range m {
res = append(res, i)
	}
	
sort.Strings(res)

	return res
}
