package main

import "testing"

type TestWord struct {
	word string

	result bool
}

func TestWorkOk(t *testing.T) {

	words := []TestWord{
		{"abcde", true},
		{"abcda", false},
	}

	for _, word := range words {
		resWord := WordOk(word.word)
		if word.result != resWord {
			t.Fatal("errroorr")
		}

	}
}

type TestRevWord struct {
	word string
	drow string
}

func TestReverseWord(t *testing.T) {

	tests := []TestRevWord{
		{"abcde", "edcba"},
	}

	for _, test := range tests {

		res := ReverseWord(test.word)
		if test.drow != res {
			t.Fatal("oh error")
		}
	}

}
