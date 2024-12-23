package main

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
)


func BenchmarkIndex(b *testing.B) {
	for _, length := range []int{10, 100, 1000, 10000} {
		phrase := randomPhrase(length)
		name := fmt.Sprintf("Allocate-%d", length)
		b.Run(name, func(b *testing.B) {
			
			for n := 0; n < b.N; n++ {
				w := MakeWords(phrase)
				w.Index(name)
			}
		})
	}
}

// randomPhrase возвращает фразу из n случайных слов.
func randomPhrase(n int) string {
	words := make([]string, n)
	for i := range words {
		words[i] = randomWord(3)
	}
	return strings.Join(words, " ")
}

var rnd = rand.New(rand.NewSource(0))

// randomWord возвращает слово из n случайных букв.
func randomWord(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyz"
	chars := make([]byte, n)
	for i := range chars {
		chars[i] = letters[rnd.Intn(len(letters))]
	}
	return string(chars)
}