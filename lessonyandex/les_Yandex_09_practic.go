package lessonyandex

import "slices"

//______________________________________________________________________________________
/*
Сестра Гоши Маша любит играть в прятки. Маша создала массив из n чисел от 1 до n.
И какие-то из них поменяла на другие, также от 1 до n.

	Напишите функцию FindMissingValues(nums []int) []int,
	 которая принимает данный массив и возвращает массив,
	  содержащий числа, которые пропали, в порядке возрастания.

Например, FindMissingValues([4, 3, 2, 7, 8, 2, 3, 1]) должна возвращать [5 6],
FindMissingValues([4, 4, 4, 4]) — [1 2 3], а FindMissingValues([1, 1, 1]) — [2 3].
*/
func FindMissingValues(nums []int) []int {
	var result []int
	m := make(map[int]int)
	slices.Sort(nums)
	for key, vol := range nums {
		m[vol] = key
	}
	for i := 1; i < len(nums)+1; i++ {
		_, ok := m[i]
		if !ok {
			result = append(result, i)
		}
	}
	return result
}

//_________________________________________________________________________________

/*
Дан массив из n целых чисел. Каждое из чисел повторяется, кроме одного.
	Найдите и выведите единственное число. Напишите функцию FindValue(nums []int) int,
которая принимает данный массив и возвращает единственное число, которое не повторяется,
*/
func FindValue(nums []int) int {
	var result int = nums[len(nums)-1]
	for i := 0; i < len(nums)-1; i++ {
		a := nums[i]
		sl := nums[i+1:]
		ok := slices.Contains(sl, a)
		if a == 0 {
			continue
		}
		if ok {
			sl[slices.Index(sl, a)] = 0
			nums[i] = 0
			continue
		} else {
			result = a
			break
		}
	}
	return result
}
//_____________________________________________________________________________________
/*
Дан массив целых чисел и два числа m строк и n столбцов. 
Необходимо конвертировать его в двумерный массив с указанным количеством строк и
 столбцов используя все числа исходного массива.
  Напишите функцию Convert2D(nums []int, m, n int) [][]int,
   которая принимает данный массив и возвращает двумерный массив.
Например, Convert2D([1, 2, 3, 4, 5, 6, 7, 8], 4, 2) должна возвращать [[1 2] [3 4] [5 6] [7 8]].
*/
func Convert2D(nums []int, m, n int) [][]int{
	result := [][]int{}
	for i:= 0; i< len(nums); i+=n {
		result = append(result, nums[i:i+n])
	}
	
	return result
	}
//______________________________________________________________________________________________
/*
Дан массив из n целых чисел. Переместите все нули в конец массива, 
при этом остальные числа должны остаться в том же порядке,
 в котором были. Напишите функцию MoveZeroes(nums []int) []int,
 которая принимает данный массив и возвращает массив с перемещёнными нулями.
*/
func MoveZeroes(nums []int) []int {
	result := []int {}
	promsl := []int {}
	for i:=0; i<len(nums); i++ {
		if nums[i] != 0{
			result = append(result, nums[i])
		}else{
			promsl = append(promsl, nums[i])
		}
	}
	result = append(result, promsl...)
	
	return result
	}
//__________________________________________________________________________________________________	


