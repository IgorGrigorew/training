package main

import (
	"fmt"
	//"strconv"
)


//счастливый номер

func main() {

	y()

}

func y() {
	var a string = "1422"
	// fmt.Scan(&a)

	byt := []byte(a)

	
	sl := slise(byt)
	
	le, re := compare(sl)

	def := le - re
	
	var res []byte

	 if def > 0 {
		res = modific(sl, def)
	} else {
		res = negative(byt)
	}

	input(res)
	
	

}

func input(sl []byte) {

	for _, n := range sl {
		fmt.Print(n)
	}

}

func slise(a []byte) []byte {

	for i, v := range a {
		a[i] = v - 48
	}
	return a
}

// изменение
func modific(a []byte, b int) []byte {

	l := len(a)
	
	for i := 1; ; i++ {
		//	fmt.Println(l,a[l-i] )
		if b+int(a[l-i]) <= 9 {

			a[l-i] += byte(b)
			break
		} else {
			b -= 9 - int(a[l-i])
			a[l-i] = 9

			continue
		}
	}
	return a
}

// сравнение
func compare(sl []byte) (int, int) {
	le, re := 0, 0
	l := len(sl)
	for left, rigt := 0, l-1; rigt > left; {
		le += int(sl[left])
		re += int(sl[rigt])
		left++
		rigt--

	}

	

	return le, re
}

func negative(a []byte) []byte {
	// fmt.Println("x", a )
	var res []byte
	//y := 1
	for i := len(a)-1; i >= 0 ; i-- {

		a[i] = 0
		if i > 0 && a[i-1] < 9 {
			a[i-1]+=1
		} else if i > 0 && a[i-1] >= 9 {

continue
		}
	//	fmt.Println(i,a)
		
		le, re := compare(a)

		if le == 0 && re == 0{
			a[len(a)-1] = 1
			res = zero(a)	
			return res
		}
//fmt.Println("xx", le, re )
		def := le - re
		

		 if def > 0 {
			res = modific(a, def)
			return res
		} else {
				continue
			}
		
		}
		return res
	}

func zero(a []byte) []byte{
a[len(a)/2 - 1] = a[len(a)-1]
	return a 
}
