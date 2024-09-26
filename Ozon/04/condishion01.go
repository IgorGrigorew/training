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
	for _, vol := range str {
		min := 15
		max := 30
		for _, v := range vol {
			zn := v[:2]
			n := v[3:]
			num, err := strconv.Atoi(n)
			if err != nil {
				fmt.Println(err)
			}
			switch zn {
			case ">=":
				if num > min {
					min = num
				}

			case "<=":
				if num < max {
					max = num
				}

			}
			if min > max {
				fmt.Println(-1)
			}else{
				fmt.Println(min)
			}
		}
fmt.Println(" ")
	}

	

}

// ____________________________________________________________________________
// Читает данные
func input() [][]string {

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

	str := make([][]string, num)

	for i := 0; i < num; i++ {

		n, err := read.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}
		nn, err := strconv.Atoi(strings.TrimSpace(n))
		if err != nil {
			fmt.Println(err)
		}
		sl := make([]string, nn)
		for y := 0; y < nn; y++ {
			s, err := read.ReadString('\n')
			if err != nil && err != io.EOF {
				fmt.Println(err)
			}
			sl[y] = strings.TrimSpace(s)

		}
		str[i] = sl

	}
	return str
}
