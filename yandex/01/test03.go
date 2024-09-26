package main

import (
	"fmt"
)
/*
неверный ответ. Подробнее 👉
Different number of lines: out = 2, corr = 5
*/
const (
	place  = 0
	target = 4
)

type Grasshopper struct {
	P int
	T int
} // знает своё местоположение на линейке

type Jumper interface {
	WhereAmI() int      // выводит текущее положение кузнечика на линейке
	Jump() (int, error) // кузнечик прыгает к зерну. Выводит новое положение кузнечика, или ошибку, если он уже ест зерно
}

func (g *Grasshopper) WhereAmI() int {
	return g.P
}

func (g *Grasshopper) Jump() (int, error) {
	tar := g.T
	pl := g.P
	fin := tar - pl
	if fin == 0 {
		return pl, fmt.Errorf("вкусное зернышко")
	} 
		if fin < 5 {
			pl += fin
		} else {
			pl += 5
		}
	
	g.P =  pl
	return  pl, nil
}

func PlaceJumper(place, target int) Jumper {
	gr := Grasshopper{
		P: place,
		T: target,
	}
	return Jumper(&gr)
}

func main() {
	g := PlaceJumper(place, target)
	fmt.Println(g.WhereAmI())
	for {
		currPlace, err := g.Jump()
		if err != nil {
			break
		}
		fmt.Println(currPlace)
	}
}
