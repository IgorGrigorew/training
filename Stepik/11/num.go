package main

import "fmt"

type slnum struct {
	sl []int
	num int
}

func (s *slnum) vol() int {
	res := 0
fn := func() func(int)int {
	res = s.sl[s.num]
	
return	func(int)int{
	s.num++
	//fmt.Println("i=",i)
		return s.num
	}
	}()
fn(s.num)

	return res
}

type intr interface {
	vol() int
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	var in intr
	str := slnum{nums, 0}
	in = &str

	
	fmt.Println(in.vol())
	fmt.Println(in.vol())
}