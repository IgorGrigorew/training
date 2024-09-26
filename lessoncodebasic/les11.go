package lessoncodebasic

// принимает вариативный список слайсов чисел и объединяет их в 1, сохраняя последовательность
func MergeNumberLists(numberLists ...[]int) []int {
	res := []int{}
	for _, i := range numberLists {
		res = append(res, i...)
	}
	return res
}
