package main

import (
	"fmt"
	"math/rand"
)

// начало решения

// генерит случайные слова из 5 букв
// с помощью randomWord(5)
func generate(cancel <-chan struct{}) <-chan string {
	out := make(chan string)
	//	defer fmt.Println("start generate")

	go func() {
		// ...
		//	defer fmt.Println("stop gorutine generate")
		defer close(out)

		for {

			res := randomWord(5)
			//	fmt.Println(res)
			select {
			case <-cancel:
				return
			case out <- res:
			}

		}

	}()

	return out
}

// выбирает слова, в которых не повторяются буквы,
// abcde - подходит
// abcda - не подходит
func takeUnique(cancel <-chan struct{}, in <-chan string) <-chan string {
	out := make(chan string)

	go func() {
		// ...
		//defer fmt.Println("stop takeUnique")
		defer close(out)

		for word := range in {

			if WordOk(word) {
				select {
				case <-cancel:

					return

				case out <- word:

				}

			}

		}
	}()

	return out
}

func WordOk(word string) bool {
	m := make(map[rune]struct{})

	for _, vol := range word {
		_, ok := m[vol]
		if ok {
			return false
		} else {
			m[vol] = struct{}{}
		}

	}

	return true
}

// переворачивает слова
// abcde -> edcba
func reverse(cancel <-chan struct{}, in <-chan string) <-chan string {
	out := make(chan string)

	go func() {
		// ...
		//	defer fmt.Println("stop reverse")
		defer close(out)

		for word := range in {

			select {

			case <-cancel:
				return
			case out <- ReverseWord(word):

			}
		}
	}()

	return out
}

func ReverseWord(word string) string {
	l := len(word) - 1
	byt := []byte(word)

	for i, y := 0, l; i <= y; i++ {
		byt[i], byt[y] = byt[y], byt[i]
		y--
	}
	return word + "->" + string(byt)
}

// объединяет c1 и c2 в общий канал
func merge(cancel <-chan struct{}, c1, c2 <-chan string) <-chan string {
	out := make(chan string)

	go func() {
		//	defer fmt.Println("end merge")
		defer close(out)

		for {
			select {
		case vol1, ok1 := <-c1:
			if ok1 {
			select {
			case <-cancel:
				return
			case out <- vol1:

			}
			}
			

		case vol2, ok2 := <-c2:
			if ok2 {
				select {
			case <-cancel:
				return
			case out <- vol2:

			}
			}
			
		case <-cancel:
			return

		}
		}
		

	}()
	return out
}

// печатает первые n результатов
func print(cancel <-chan struct{}, in <-chan string, n int) {
	// ...
	go func() {
		
		for i:=0; i < n; i++{
			
			select {
			case <- cancel:
				return
			case vol, ok := <- in:
				if !ok {
					return
				}
				fmt.Println(vol)
			}

		}

	}()

}

// конец решения

// генерит случайное слово из n букв
func randomWord(n int) string {
	const letters = "aeiourtnsl"
	chars := make([]byte, n)
	for i := range chars {
		chars[i] = letters[rand.Intn(len(letters))]
	}
	return string(chars)
}

func main() {
	cancel := make(chan struct{})
	defer close(cancel)

	c1 := generate(cancel)
	c2 := takeUnique(cancel, c1)
	c3_1 := reverse(cancel, c2)
	c3_2 := reverse(cancel, c2)
	c4 := merge(cancel, c3_1, c3_2)
	print(cancel, c4, 10)
	fmt.Println("stop")
}
