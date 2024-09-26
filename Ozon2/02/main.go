package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

/*
 первое число кода вершины — её номер 𝑢u,

 второе число кода вершины — количество её детей 𝑑𝑢 du, называемое также степенью этой вершины,

 затем следуют 𝑑d чисел — номера сыновей в произвольном порядке.
*/


func main() {

	y()

}

func y() {

	file, err := os.Open("text.txt")
	if err != nil {
		fmt.Print(err)
	}
defer file.Close()
	
	var in *bufio.Reader
	var out *bufio.Writer
	/*
		in = bufio.NewReader(os.Stdin)

	*/

	in = bufio.NewReader(file)

	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	st, err := in.ReadString('\n')
		if err != nil {
			fmt.Print(err)
		}

	num, err := strconv.Atoi(strings.TrimSpace(st))
if err != nil {
	fmt.Print(err)
}
	
 

	for i := 0; i < num; i++ {

	

	str1, err := in.ReadString('\n')
		if err != nil {
			fmt.Print("b",err)
		}


		str2, _ := in.ReadString('\n')
		


	 Stringconv( str1, str2)

	
	
//fmt.Fprintln(out, "x", str1)
//fmt.Fprintln(out, "y", str2)


	}


}


func Stringconv( a, b string) {

	n, err := strconv.Atoi(strings.TrimSpace(a))
if err != nil {
	fmt.Print(err)
}

sl :=[]int32{}



for _, v := range b{
if v < 47 {
	continue
}
	sl =append(sl, int32(v)- 48) 
	
}

//fmt.Println(sl)



	 Modific( n, sl...)


}


func Modific( n int, sl ...int32)  {

	sltop := []int32{}
	 slchil := []int32{}
	i:= 0
	for i < n {
	
	sltop = append(sltop, sl[i])
	
		i++
		if sl[i] == 0 {
			i++
			continue
		}
		chl:=sl[i]
		i++
	slchil = append(slchil, sl[i: i + int(chl)]...)
	i+=int(chl)
	continue
	
	}

 slices.Sort(sltop)
 slices.Sort(slchil)
	fmt.Println("top", sltop)
	fmt.Println("son", slchil)


for i:= 0 ; i<len(sltop)-1; i++ {
if sltop[i] == slchil[i]{
	continue
	
}else{
	fmt.Println(sltop[i])
	break

}

}





	/*
	первое число кода вершины — её номер 
	
	 второе число кода вершины — количество её детей 
	
	, называемое также степенью этой вершины,
	
	 затем следуют 
	
	 чисел — номера сыновей в произвольном порядке.
	*/
	
	
	
	



}