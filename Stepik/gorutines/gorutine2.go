package main

import (
	"fmt"
	"strings"
	"unicode"
)

// counter хранит количество цифр в каждом слове.
// ключ карты - слово, а значение - количество цифр в слове.
type counter map[string]int

// countDigitsInWords считает количество цифр в словах фразы
func countDigitsInWords(phrase string) counter {
	words := strings.Fields(phrase)
	counted := make(chan int)

	// начало решения

stats := make(counter)

for _, w := range words {

	go func() {

		counted <- countDigits(w)

		// Пройдите по словам,
		// посчитайте количество цифр в каждом,
		// и запишите его в канал counted
	}()

	num := <- counted
	// Считайте значения из канала counted
	// и заполните stats.

	stats[w] = num
	// В результате stats должна содержать слова
	// и количество цифр в каждом.
}
	// конец решения

	return stats
}

// countDigits возвращает количество цифр в строке
func countDigits(str string) int {
	count := 0
	for _, char := range str {
		if unicode.IsDigit(char) {
			count++
		}
	}
	return count
}

// printStats печатает слова и количество цифр в каждом
func printStats(stats counter) {
	for word, count := range stats {
		fmt.Printf("%s: %d\n", word, count)
	}
}

func main() {
	phrase := "0ne 1wo thr33 4068"
	stats := countDigitsInWords(phrase)
	printStats(stats)
}
