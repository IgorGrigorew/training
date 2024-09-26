package stepik

import (
	"container/list"
	"fmt"
	//"math/rand"
	//"sync"
//	"time"
)

/*
Реализуйте функции для очереди FIFO (First In, First Out (ПЕРВЫЙ пришел, ПЕРВЫЙ вышел)) с помощью списков.
Должны быть данные функции:

Push (добавление элемента)
Pop (удаление элемента и его возврат)
printQueue (печать очереди в одну строку без пробелов)
Функцию main писать не нужно! Писать код вне функций не нужно.
*/

func Push(elem interface{}, queue *list.List) {
	// ...
	queue.PushBack(elem)
}

func Pop(queue *list.List) interface{} {
	// ...
	queue.Remove(queue.Front())
	return nil
}

func printQueue(queue *list.List) {
	// ...
	for temp := queue.Front(); temp != nil; temp = temp.Next() {
		fmt.Print(temp.Value)
	}
}


func main4() {
	queue := list.New()

    Push(1, queue)

    Push(2, queue)

    Push(3, queue)

    printQueue(queue) // 123

    Pop(queue)

    printQueue(queue) // в ту же строку: 23

}
	