package lessonyandex

import (
	"bytes"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

/*
Напишите функцию GetCharacterAtPosition(str string, position int) (rune, error) для робота-помощника,
 которая получает на вход строку и целое число. Функция должна возвращать символ строки,
который находится на позиции, указанной пользователем (и nil в качестве ошибки).
Если пользователь ввёл число, которое выходит за пределы длины строки,
 функция должна возвращать в качестве ответа нулевую руну (0) сообщение об ошибке (position out of range).
*/
func GetCharacterAtPosition(str string, position int) (rune, error){
	var res rune
if position > len([]rune(str)){
	return 0, fmt.Errorf("position out of range")
}
  for i, vol := range str {
	if position == i {
		res = vol
	}
  } 
return res, nil
}
//________________________________________________________________________________________
/*
Напишите функцию Factorial(n int) (int, error) для нахождения факториала без ошибок. 
Функция получает на вход целое число и выводит его факториал. Если пользователь ввёл отрицательное число, 
программа должна вернуть сообщение об ошибке (factorial is not defined for negative numbers).
*/
func Factorial(n int) (int, error){
	if n <0 {
		return 0, fmt.Errorf("factorial is not defined for negative numbers")
	}
	return Fac(n), nil
	}
	
	func Fac(n int) int  {
		if n == 0 {
			return 1
		}
		return n* Fac(n-1)
	}
//_______________________________________________________________________________________________
/*
Напишите функцию IntToBinary(num int) (string, error),
которая принимает на вход целое число и возвращает его двоичное представление.
Если пользователь ввёл отрицательное число,
 программа должна возвращать сообщение об ошибке (negative numbers are not allowed).
*/
func IntToBinary(num int) (string, error) {
	if num < 0 {
		return "", fmt.Errorf("negative numbers are not allowed")
	}
	result := strconv.FormatInt(int64(num), 2)
	return result, nil
	}
//______________________________________________________________________________________________-	
	/*
Напишите функцию SumTwoIntegers(a, b string) (int, error), которая получает две строки и,
если это целые числа, возвращает их сумму. Если пользователь ввёл данные не целого типа,
функция должна вернуть сообщение об ошибке (invalid input, please provide two integers).
*/
func SumTwoIntegers (a, b string)(int, error) {
	var result int
if s1, err := strconv.Atoi(a); err != nil {
return 0, fmt.Errorf("invalid input, please provide two integers")
}else{
	result+=s1
}
if s2, err := strconv.Atoi(b); err != nil {
	return 0, fmt.Errorf("invalid input, please provide two integers")
	}else{
		result+=s2
	}
return result, nil
}
//________________________________________________________________________________________________



func ConcatStringsAndInt(str1, str2 string, num int) string{

	return fmt.Sprintf("%s%s%d", str1, str2, num)
	
	}
	
	//___________________________________________________
	
	func DivideIntegers(a, b int) (float64, error){
		if b == 0{
			return 0, fmt.Errorf("division by zero is not allowed")
		}
		return float64(a) / float64(b), nil
	}
	
	//______________________________________________________
	
	func GetCharacterAtPosition1(str string, position int) (rune, error){
	s :=  []rune(str)
	if len(s) < position {
		return 0, fmt.Errorf("position out of range")
	}
	
	
	return  s[position], nil
	}
	
	//___________________________________________________________
	
	func AreAnagrams(str1, str2 string) bool{
	s1 := strings.ToLower(str1)
	s2 := strings.ToLower(str2)
		
	byt1 := []byte(s1)
		slices.Sort(byt1)
	
		byt2 := []byte(s2)
		slices.Sort(byt2)
	return bytes.Equal(byt1, byt2)
	
	
	
	}

