package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)



func main() {
treatment()

}


//обрабатывает
func treatment() {
	
str := input()
const layout = "2.1.2006"
for _, vol := range str{
	
	t := strings.ReplaceAll(vol, " ", ".")
	_ , err := time.Parse(layout, t)

	if err != nil{
		fmt.Println( "No")
	}else{
		fmt.Println("Yes")
	}
}

}


//____________________________________________________________________________
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
	str[i] = strings.TrimSpace(s)
	
}
return str
}