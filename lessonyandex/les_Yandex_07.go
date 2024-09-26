package lessonyandex

import (
	"cmp"
	"slices"
	"strings"
)

/* Напишите функцию FindMaxKey(m map[int]int) int,
которая принимает мапу и возвращает значение наибольшего ключа. */
 func FindMaxKey(m map[int]int) int {
	var res int = -1000
	for key := range m {
if key > res {
	res = key
}
	}
	return res
}

/*Напишите функцию SumOfValuesInMap(m map[int]int) int,
 которая возвращает сумму значений в мапе. */
func SumOfValuesInMap(m map[int]int) int {
	var res int 
	for _, vol := range m {
res += vol
	}
	return res
}

/* Напишите функцию SwapKeysAndValues(m map[string]string) map[string]string,
 которая принимает мапу и возвращает новую мапу, где ключи и значения поменялись местами. */
func SwapKeysAndValues(m map[string]string) map[string]string {
result := make(map[string]string)
for key, vol := range m {
	result[vol] = key
}
return result
}

/* Палиндром - строка, которая читается одинаково слева направо и справа налево.
 Напишите функцию IsPalindrome(input string) bool, 
 которая принимает строку и проверяет, является ли она палиндромом.*/
 func IsPalindrome1(input string) bool {
	input =  strings.ToLower(strings.ReplaceAll(input, " ", ""))
	runinp := []rune (input)
	for i, vol := range runinp {
		if vol != runinp[len(runinp)-(i+1)] {
			return false
				}
	}
	return true
	}

	/* Напишите функцию Permutations(input string) []string,
которая принимает строку и выводит все перестановки её символов в алфавитном порядке.
С len() >= 4 работает не правильно */
func Permutations1(input string) []string {
	result := []string{}
   //var resultstr string
   for i:= 0; i < len(input); i++ {
	   var str, revstr, prom  string
	   if i > 0  {
	   for y := 1;  ; y++ {                   // y:= i  ==> y:=1
		   if len(prom) == len(input){
			   break
		   }
		   if y == len(input)  {            // -1
		   y = 0
		   }
		   prom += string(input[y])
	   }
	   input = prom
		   }
	   
	   for i, vol := range (input) {
		   str += string(vol)
		   revstr += string(input[len(input)-(i+1)])
	   }
   result = append(result, str, revstr)
   // выход из цикла
   if len(input) == 2 {
	   break
   }
   }
   slices.SortFunc(result, func(a, b string) int {
	   return cmp.Compare(strings.ToLower(a), strings.ToLower(b))
   })
   return  result
   }
   