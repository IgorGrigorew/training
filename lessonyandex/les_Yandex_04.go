package lessonyandex

import (
	
	"math"
)

/* Напишите функцию FindIntersection(k1, b1, k2, b2 float64) (float64, float64),
которая принимает на вход четыре вещественных числа k1, b1, k2 и b2,
представляющих коэффициенты уравнений y = kx + b для двух различных астрономических объектов.
 Функция возвращает координаты точки, в которой данные астрономические траектории пересекаются.
Если траектории параллельны, программа должна вывести math.NaN().*/
func FindIntersection(k1, b1, k2, b2 float64) (float64, float64) {
	var x float64
	var y float64
	if k1 == k2 && b1 != b2 {
		return math.NaN(), math.NaN()
	} else if k1 != k2 {
		x =  (-b1 + b2)/ (k1 + (-k2)) 
		y = k1*x + b1
	}
	
return x, y
}
