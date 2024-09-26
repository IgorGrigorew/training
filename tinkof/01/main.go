package main

import (
	"bufio"
	"fmt"
    "io"
	"os"
//	"strings"
	
//	"strconv"
)



const (MON = iota+1
 	TUE 
 	WED
 	THU
 	FRI
 	SAT
 	 SUN)

 func input() []int {
	
file, err := os.Open("text.txt")
if err != nil {
    fmt.Println(err)
}
defer file.Close()


	read := bufio.NewReader(file)
	//read := bufio.NewReader(os.Stdin)

 sl := []int{}

for i:=1; ; i++ {


    s, err := read.ReadString(' ')
	if err != nil && err != io.EOF {
       fmt.Println(err)
	   break
    }
	switch s {
	case "MON" :sl = append(sl, MON+i)
	case "TUE" :sl = append(sl, TUE+i)
	case "WED" :sl = append(sl, WED+i)
	case "THU" :sl = append(sl, THU+i)
	case "FRI" :sl = append(sl, FRI+i)
	case "SAT" :sl = append(sl, SAT+i)
	case "SUN" :sl = append(sl, SUN+i)
	case " " : sl = append(sl, 0)
	
	}
	

}
return sl
}

func t()  {
sl := input()




fmt.Print(sl)


	}



func main() {
t()
}