package lessoncodebasic

import (
	"strings"
)

// func Map(strs []string, mapFunc func(s string) string) []string,
//
//	преобразует каждый элемент слайса strs с помощью функции mapFunc и возвращает новый слайс.
//
// Учтите, что исходный слайс, который передается как strs, не должен измениться в процессе выполнения.
func Map(strs []string, mapFunc func(s string) string) []string {
	srez := make([]string, len(strs))
	//fmt.Println("срез в футкции Map", srez, strs)
	for i, vol := range strs {
		srez[i] = mapFunc(vol)
	}
	return srez
}

// изменяет строки
func MapFunc(s string) string {
	//fmt.Println("строка в функции mapFunc", s)
	if strings.ToLower(s) == s {
		return strings.ToUpper(s)
	} else {
		return strings.ToLower(s)
	}
}
