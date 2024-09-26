package lessonyandex

// рекурсивно вычисляет сумму цифр числа
func SumDigitsRecursive(n int) int {
	if n < 10 {
		return n
	}
	//return n%10 + n/10
	return SumDigitsRecursive(n%10) + SumDigitsRecursive(n/10)
}
