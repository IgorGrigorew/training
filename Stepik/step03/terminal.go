package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// позиция курсора
type position struct {
	x int
	y int
}

// перемещение вверх
func (p *position) Up() {
	if p.y > 0 {
	p.y -= 1	
	}
}

// перемещение вниз
func (p *position) Down() {
	
	p.y += 1	

}

// перемещение влево
func (p *position) Left() {
	if p.x > 0 {
	p.y -= 1	
	}
}

// перемещение вправо
func (p *position) Rigt(sl [][]string) {
	if p.y < len(sl[p.x])-1 {
	p.y += 1	
	}
}

// перемещение в начало строки
func (p *position) Home() {
	
	p.y = 0	
	
}


// перемещение в конец строки
func (p *position) End(sl [][]string) {
	
	p.y = len(sl[p.x])-1	
	
}

// перемещение на следующюю строку
func (p *position) Enter(sl [][]string) {
	
	if p.y == len(sl[p.x])-1 {

	}	
	
}

// ввод
func (p *position) Input(s string, sl [][]string) {
	sl[p.x][p.y] = s
	p.y ++
}


func main() {
	

	treatment()

}

// обрабатывает
func treatment() {

	str := input()







fmt.Println(str)

}


// ____________________________________________________________________________
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
	if err != nil {
		fmt.Println(err)
	}

	num, err := strconv.Atoi(strings.TrimSpace(s))
	if err != nil {
		fmt.Println(err)
	}

	str := make([]string, num)

	for i := 0; i < num; i++ {


		s, err := read.ReadString('\n')
		if err != nil && err != io.EOF {
			fmt.Println(err)
		}
		str[i] = strings.TrimSpace(s)

	}

	return str
}
