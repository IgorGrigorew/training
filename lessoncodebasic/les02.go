package lessoncodebasic

/*Реализуйте функцию IntsCopy(src []int, maxLen int) []int, которая создает копию слайса src с длиной maxLen.
Если maxLen равен нулю или отрицательный, то функция возвращает пустой слайс []int{}.
Если maxLen больше длины src, то возвращается полная копия src. */

func IntsCopy(src []int, maxLen int) []int {
	if maxLen < 1 {
		return []int{}
	} else if maxLen > len(src) {
		return src
	}
	copnum := make([]int, maxLen)
	copy(copnum, src)
	return copnum
}
