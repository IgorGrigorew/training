package main

import (
//	"fmt"
)

// начало решения

// Avg - накопительное среднее значение.
type Avg [T float64 | int] struct{
	valAvg T
	num T
}

// Add пересчитывает среднее значение с учетом val.
func (a *Avg[T]) Add(val T) *Avg[T] {
	// ...
	
a.valAvg += val
a.num ++

return a
}

// Val возвращает текущее среднее значение.
func (a Avg[T]) Val() T {
	// ...
	if a.num == 0{
		return a.valAvg
	}
	return a.valAvg/a.num
}

// конец решения
/*
func main() {
	intAvg := Avg[int]{}
	intAvg.Add(4).Add(3).Add(2)
	fmt.Println(intAvg.Val()) // 3

	floatAvg := Avg[float64]{}
	floatAvg.Add(4.0).Add(3.0)
	floatAvg.Add(2.0).Add(1.0)
	fmt.Println(floatAvg.Val()) // 2.5
}
*/