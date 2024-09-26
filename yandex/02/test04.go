package main

import (
	"bufio"
	"fmt"
	"strings"

	//"io"
	"os"
	"strconv"
	//"strings"
)
/*
На первой строке передается количество людей в списке - натуральное число.
 Затем на следующих строках идет список людей. А после списка - перечень префиксов всех тех, кого надо найти.
  Перечень префиксов может быть произвольной длины. Пустая строка вместо префикса значит что префиксы закончились.
*/
func main() {

	file, err := os.Open("text.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()	
	

	name := []string{}
	pref := []string{}

	i, a :=0, 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if i == 0 {
			n, err := strconv.Atoi(line)
			a = n
			if err != nil {
				fmt.Println(err)
			}
		} else if i>0 && i <=a {
			name = append(name, line)
		} else {
			pref = append(pref, line)
		}
		i++
    }

for _, pr := range pref {
	ok := false
	for _, nm := range name {
		if strings.HasPrefix(nm, pr) {
			ok = true
			fmt.Println(nm)
			break
		}
	}
	if !ok {
		fmt.Println("Не найдено")
	}
}

	
	
//fmt.Println(name)
//fmt.Println(pref)


}