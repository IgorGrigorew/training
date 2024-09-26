package main

import "fmt"

func main() {
	var a, b int = 3,3
	
//	fmt.Scan(&a, &b)
	for i := a; i > 0; i-- {
		var res string
		for y := 1; y <= i; y++ {
			res += "*"
		}
		fmt.Println(res)
	}

	i:=0
	if a != 0 {
		i++
	}
	for ; i < b; i++ {
		fmt.Println("*")
	}
}
