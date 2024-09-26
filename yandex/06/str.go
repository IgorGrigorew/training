package main

import (
	"bufio"
//	"slices"
	"time"

	"fmt"
	"os"
	//"strconv"
)



func main() {
	start := time.Now()
	fmt.Println("Start")
	yandex()
	end := time.Now()

	fmt.Println("Stop", end.Sub(start))
}

func yandex() {

a, b, c := input()

for i := 0; ; i ++ {

	switch {
	case  i == 0:
		if a[i] != b[i] || b[i] != c[i]{
	
		fmt.Println("IMPOSSIBLE")
		break
	}
case i == 0:

}


	



}


}



func input() (string, string, string){

	file, err := os.Open("text.txt")
	if err != nil {
		fmt.Print(err)
	}

	var in *bufio.Reader
	
	//	in = bufio.NewReader(os.Stdin)

	in = bufio.NewReader(file)

	

	var  a, b, c string
	

	
	fmt.Fscan(in, &a, &b, &c)
	

	return a, b, c
}
