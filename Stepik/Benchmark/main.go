package main

// не удаляйте импорты, они используются при проверке
import (
	
	// "os"
	"fmt"
	"strings"
)

// ┌─────────────────────────────────┐
// │ не меняйте код ниже этой строки │
// └─────────────────────────────────┘

// Библиотечная
func MatchContains(pattern string, src string) bool {
    return strings.Contains(src, pattern)
}

// Самописная
func MatchContainsCustom(pattern string, src string) bool {
    if pattern == "" {
        return true
    }
    if len(pattern) > len(src) {
        return false
    }
    pat_len := len(pattern)
    for idx := 0; idx < len(src)-pat_len+1; idx++ {
        if src[idx:idx+pat_len] == pattern {
            return true
        }
    }
    return false
}


func main() {
	s := "go is awesome"
	fmt.Println(MatchContains("is", s))
	// true
	fmt.Println(MatchContains("go.*awesome", s))
	// false

	fmt.Println(MatchContainsCustom("go.*awesome", s))
	// true
	fmt.Println(MatchContainsCustom("^go", s))
	// true
	fmt.Println(MatchContainsCustom("awesome$", s))
	// true
}