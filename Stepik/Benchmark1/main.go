package main

// не удаляйте импорты, они используются при проверке
import (
    "fmt"
   // "math/rand"
  //  "os"
  //  "testing"
)

// IntSet реализует множество целых чисел
// (элементы множества уникальны).
type IntSet struct {
	Map map[int]struct{}
}

// MakeIntSet создает пустое множество.
func MakeIntSet() IntSet {
	m := make(map[int]struct{}, 0)
    return IntSet{m}
}

// Contains проверяет, содержится ли элемент в множестве.
func (s IntSet) Contains(elem int) bool {
  _, ok := s.Map[elem]

	return ok
}

// Add добавляет элемент в множество.
// Возвращает true, если элемент добавлен,
// иначе false (если элемент уже содержится в множестве).
func (s IntSet) Add(elem int) bool {
   
	if s.Contains(elem) {
return false
	}
	s.Map[elem] = struct{}{}
	return true
}

func main(){

	m := MakeIntSet()

	m.Add(5)

	fmt.Println(m)

}