package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	

	treatment()

}

// обрабатывает
func treatment() {

	str := input()
//	slint := make([][]int, len(str))
	for _, vol := range str {

	 res := slise(convert(vol))
 	fmt.Println(len(res))
	var st string
		for _, s := range res {
			s1 := strconv.Itoa(s)
	st += s1 + " "
		}
	 fmt.Println(st)

	}

}

func slise (sl []int) [] int {
	if len(sl) < 2 {
		sl = append(sl, 0)
		return sl
	}
		compres := 0
		start := sl[0]
result := []int{}
nap := 0
	for i := 1; i < len(sl); i++ {
		n, ok := fn(sl[i-1], sl[i])
//fmt.Println("resultat", n, ok)
//fmt.Println("dano", sl[i-1], sl[i])
		
			
		if ok {
			if nap != n && nap != 0 {
				result = append(result, start, compres)
				compres = 0
				start = sl[i]
				nap = 0
				if i == len(sl)-1 {
					result = append(result, start, compres)
				}
				//fmt.Println("feil")
				continue
			}
			compres += n
			if i == len(sl)-1 {
				result = append(result, start, compres)
			}
			nap = n
		}else{
			result = append(result, start, compres)
			compres = 0
			start = sl[i]
			nap = 0
			if i == len(sl)-1 {
				result = append(result, start, compres)
			}
		}


	}

return result
}

func fn (a , b int) (int, bool) {
	res := b - a
	if res == -1 {
		return res, true
	} else if res == 1 {
		return res, true
	}

	return 0, false
}

func convert(s string) []int {
	result := []int{}
	min := 0
	for i, r := range s {

		if r == ' ' || i == len(s) {
			n, err := strconv.Atoi(s[min:i])
			if err != nil {
				fmt.Println(err)
			}
			result = append(result, n)
			min = i + 1
		}

		if i == len(s)-1 {
			n, err := strconv.Atoi(s[min:])
			if err != nil {
				fmt.Println(err)
			}
			result = append(result, n)
		}

	}
	return result
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

		_, err := read.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}

		s, err := read.ReadString('\n')
		if err != nil && err != io.EOF {
			fmt.Println(err)
		}
		str[i] = strings.TrimSpace(s)

	}

	return str
}