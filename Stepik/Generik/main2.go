package main

import (
//	"fmt"
)

// начало решения

// Map - карта "ключ-значение".
type Map[K comparable, V comparable] struct{
	Key []K
	Val []V
}

// Set устанавливает значение для ключа.
func (m *Map[K, V]) Set(key K, val V) {
	// ...
	m.Key = append(m.Key, key)
	m.Val = append(m.Val, val)
}

// Get возвращает значение по ключу.
func (m Map[K, V]) Get(key K) V {
	// ...
	var res V
	for i, v := range m.Key{
		if key == v {
res = m.Val[i]
		}
	}
	return res
}

// Keys возвращает срез ключей карты.
// Порядок ключей неважен, и не обязан совпадать
// с порядком значений из метода Values.
func (m Map[K, V]) Keys() []K {
	// ...
	return m.Key
}

// Values возвращает срез значений карты.
// Порядок значений неважен, и не обязан совпадать
// с порядком ключей из метода Keys.
func (m Map[K, V]) Values() []V {
	// ...
	return m.Val
}

// конец решения
/*
func main() {
	m := Map[string, int]{}
	m.Set("one", 1)
	m.Set("two", 2)

	fmt.Println(m.Get("one")) // 1
	fmt.Println(m.Get("two")) // 2

	fmt.Println(m.Keys())   // [one two]
	fmt.Println(m.Values()) // [1 2]
}
*/