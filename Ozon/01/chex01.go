package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)
const (four =  iota+1
	three 
	two
 	one) 


func main() {
treatment()

}


//обрабатывает
func treatment() {
	
str := input()

for _, vol := range str {
var on, tw, th, f int
	for _, st := range vol {
		switch st {
		case '1': on++
		case '2': tw++
		case '3': th++
		case '4': f++
		}
	}
switch{
case on != one: fmt.Println("No")
case tw != two: fmt.Println("No")
case th != three: fmt.Println("No")
case f != four: fmt.Println("No")
default: fmt.Println("Yes")
}
}


}



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
	if err != nil{
		fmt.Println(err)
	}
//fmt.Println(s, len(s))
	num, err := strconv.Atoi(strings.TrimSpace(s))
if err != nil{
	fmt.Println(err)
}

 str := make([]string, num)

for i:=0; i<num; i++ {

    s, err := read.ReadString('\n')
	if err != nil && err != io.EOF {
       fmt.Println(err)
    }
	str[i] = s
	
}
return str
}