package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"unicode"
)
/*
Вам задана строка из цифр и прописных букв латинского алфавита.
 Можно ли разделить её пробелами на последовательность корректных автомобильных номеров? 
 Иными словами, проверьте, что заданная строка может быть образована как последовательность
  корректных автомобильных номеров, которые записаны подряд без пробелов.
   В случае положительного ответа выведите любое такое разбиение.
*/
func main() {
	treatment()

}

// обрабатывает
func treatment() {

	str := input()
	for _, s := range str {

		res, ok := fn(s)
		if !ok {
			fmt.Println("-")
			continue
		}
		fmt.Println(res)
	}

}

func fn(s string) (string,  bool) {
	n, num := 0, 0
	var result string
	var ok, shortnum bool
	for i, r := range s {
		switch  {
		case i-n == 0 : if  unicode.IsLetter(r) {
			result += string(r)
			num++
			ok = false
		//	fmt.Println("case0",result, i, n, num)
			continue
		} else {
		//	fmt.Println("case0","no", string(r),result, i, n, num )
			return "-",  false
		}
		case i-n == 1 : if unicode.IsDigit(r) {
			result += string(r)
			num++
			ok = false
		//	fmt.Println("case1",result, i, n)
			continue
		}else{
		//	fmt.Println("fail", i, n)
			return "-",  false
		}
		case i-n == 2: if  unicode.IsDigit(r) {
			result += string(r)
			num++
			ok = false
		//	fmt.Println("case2",result, i, n)
			continue
		} else {
			num++
			result += string(r)
			shortnum = true
		//	fmt.Println("case2",result, i)
			continue
		}
		case i-n == 3 : if  unicode.IsLetter(r) {
			result += string(r)
			num++
			ok = false
		//	fmt.Println("case3",result, i, n)
			if shortnum {
				n = num
				result += " "
				ok = true
				shortnum = false
				continue
			}
			continue
		} else {
			return "-",  false
		}
		case i-n == 4 : if  unicode.IsLetter(r) {
			result += string(r)+" "
			num++
			n = num
			ok = true
		//	fmt.Println("case4",result, i, n)
		} else {
			return "-",  false
		}

	}

	}

	return result, ok
}

// ____________________________________________________________________________
// Читает данные
func input() []string {

	file, err := os.Open("text.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	read := bufio.NewReader(file)

	//read := bufio.NewReader(os.Stdin)
	s, err := read.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}

	num, err := strconv.Atoi(strings.TrimSpace(s))
	if err != nil {
		fmt.Println(err)
	}

	str := make([]string, num)

	for i := 0; i < num; i++ {

		s, err := read.ReadString('\n')
		if err != nil && err != io.EOF {
			fmt.Println(err)
		}
		str[i] = strings.TrimSpace(s)

	}
	return str
}
