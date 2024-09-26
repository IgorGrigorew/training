package lessonyandex

import "sort"

/*Напишите функцию ReverseSlice(slice []int) []int,
которая получает в себя слайс и возвращает новый слайс,
состоящий из элементов исходного, но в обратном порядке. */
 func ReverseSlice(slise []int) []int {
	resultSL := make([]int, 0)
for i:= len(slise); i > 0; i -- {
	resultSL = append(resultSL, slise[i-1])
}
	return resultSL
}

/*Напишите функцию SumOfSlice(slice []int) int,
 которая получает слайс из целочисленных элементов и возвращает
  сумму всех элементов слайса. */
  func SumOfSlice(slice []int) int{
	res := 0
	for _,i := range slice {
		res += i
	}
	return res
	}

	/*Напишите функцию IntersectionOfSlices(slice1, slice2 []int) []int,
 которая получает два слайса из целочисленных элементов и возвращает
  слайс с элементами, которые встречаются в обоих.*/
func IntersectionOfSlices(slise1, slise2 []int) []int {
	mapa := make(map[int]int)
   sl := []int{}
   for i, vol := range slise1 {
   mapa[vol] = i
   }
   for _, vol := range slise2 {
   _, ok := mapa[vol]
   if ok {
   sl = append(sl, vol)
   }
   }
   return sl
   }

   /*Напишите функцию SortSlice(slice []int) []int,
которая получает слайс из целочисленных элементов,
сортирует его по возрастанию и возвращает отсортированный слайс. */
func SortSlice(slice []int) []int {
	sort.Slice(slice, func(i, j int) bool {return slice[i]< slice[j]})
		return slice
	}

  //__________________________________________________________________________________________________

  /* Напишите функцию CountSubslices(slice []int) int,
 которая принимает слайс целочисленных элементов, находит все подслайсы,
 сумма элементов которых больше среднего значения суммы элементов всего слайса,
  и возвращает их количество.*/
func CountSubslices(slice []int) int {
  sred := SumSlice(slice)/ len(slice)
  var result int 
  for i := len(slice)-1; i>0 ; i-- {
    var sumint int                                      //i 2, v 3
    for y:= 0; y <=i ; y ++ {
      if sumint == sred {
        continue
      }
      result ++
      sumint += slice[i-y]
       if sumint < sred  {
        break
      }
    }
  }
  return result
  }
  // находим сумму элементов слайса
  func SumSlice(slice []int) int{
    var res int
    for _, vol := range slice {
      res += vol
    }
    return res 
  }

  //______________________________________________________________________________________________

