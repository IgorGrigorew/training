package main

import "testing"

// используйте эту переменную в бенчмарках
const src = "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."

// используйте эту переменную в бенчмарках
const pattern = "commodo"

// реализуйте бенчмарк для MatchContains
func BenchmarkMatchContains(b *testing.B) {
    // ...
	for i :=0; i < b.N; i++{
		MatchContains(pattern, src)
	}
}

// реализуйте бенчмарк для MatchContainsCustom
func BenchmarkMatchContainsCustom(b *testing.B) {
    // ...
	for i :=0; i < b.N; i++{
		MatchContainsCustom(pattern, src)
	}
}
