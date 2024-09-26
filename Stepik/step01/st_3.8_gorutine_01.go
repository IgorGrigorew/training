package stepik

import (
	"fmt"
	"sync"
	"time"
)

/*
Напишите элемент конвейера (функцию), что запоминает предыдущее значение и отправляет
 значения на следующий этап конвейера только если оно отличается от того, что пришло ранее.

Ваша функция должна принимать два канала - inputStream и outputStream,
в первый вы будете получать строки, во второй вы должны отправлять значения без повторов.
 В итоге в outputStream должны остаться значения, которые не повторяются подряд. Не забудьте закрыть канал ;)

Функция должна называться removeDuplicates()
*/
func removeDuplicates(inputStream chan string, outputStream chan string) {
	defer close(outputStream)
var  strMemory string
for  str := range  inputStream{
	if str != strMemory {
outputStream <- str
strMemory = str
//fmt.Println("func", str)
	}
	
}
}

func Print (ch chan string) {
	for  {
		vol := <-  ch
		fmt.Println(vol)
	}
	
}

func goRutine() {
	str := "1222334566678"
	 inputStream := make (chan string)
	 outputStream := make (chan string)
	go removeDuplicates(inputStream, outputStream)
	go Print(outputStream)
defer close(inputStream)
	for _, vol := range str {
		inputStream <- string(vol)
	}
	time.Sleep(time.Second*2)
}
//_______________________________________________________________________________________________

/*
Внутри функции main (функцию объявлять не нужно), вам необходимо в отдельной горутине вызвать функцию work()
 и дождаться результатов ее выполнения.

Функция work() ничего не принимает и не возвращает.
*/

func mainGorutine(){
	ch := make(chan struct{})
	
	go func (){
		defer close(ch)
	work()
	}()
	<-ch
	}
	
	func work (){
		fmt.Println("hello")
	}
//________________________________________________________________________________________

/*
Внутри функции main (функцию объявлять не нужно), вам необходимо в отдельной горутине вызвать функцию work()
 и дождаться результатов ее выполнения.

Функция work() ничего не принимает и не возвращает.
*/
func mainWaitgroup(){
	wg := new(sync.WaitGroup)
	
	for i:=0; i<=10; i++ {
		wg.Add(1)
	go func (){
		defer wg.Done()
	work1()
	}()
	}
	wg.Wait()	
	}
	
	func work1 (){	
		fmt.Println("hello")
	}
//________________________________________________________________________________
/*
Вам необходимо написать функцию calculator следующего вида:
func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int
Функция получает в качестве аргументов 3 канала, и возвращает канал типа <-chan int.
в случае, если аргумент будет получен из канала firstChan, в выходной (возвращенный)
	канал вы должны отправить квадрат аргумента.
в случае, если аргумент будет получен из канала secondChan, в выходной (возвращенный)
канал вы должны отправить результат умножения аргумента на 3.
в случае, если аргумент будет получен из канала stopChan, нужно просто завершить работу функции.
Функция calculator должна быть неблокирующей, сразу возвращая управление.
Ваша функция получит всего одно значение в один из каналов - получили значение, обработали его, завершили работу.
После завершения работы необходимо освободить ресурсы, закрыв выходной канал, если вы этого не сделаете,
	то превысите предельное время работы.
*/
func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
	result := make(chan int)

	go func() {
		defer close(result)
		select {
		case a := <-firstChan:
			result <- a * a
		case b := <-secondChan:
			result <- b * 3
		case <-stopChan:
			return
		}
	}()

	return result
}
//_______________________________________________________________________________________
/*
	Вам необходимо написать функцию calculator следующего вида:
func calculator(arguments <-chan int, done <-chan struct{}) <-chan int
В качестве аргумента эта функция получает два канала только для чтения, возвращает канал только для чтения.
Через канал arguments функция получит ряд чисел, а через канал done - сигнал о необходимости завершить работу.
	Когда сигнал о завершении работы будет получен, функция должна в выходной (возвращенный)
	канал отправить сумму полученных чисел.
Функция calculator должна быть неблокирующей, сразу возвращая управление.
Выходной канал должен быть закрыт после выполнения всех оговоренных условий, если вы этого не сделаете,
	то превысите предельное время работы.
*/
func calculator2(arguments <-chan int, done <-chan struct{}) <-chan int {
	result := make(chan int)
	go func() {
		b := 0
		defer close(result)
		for {
			select {
			case a := <-arguments:  //читаем из канала все данные
				b += a
				fmt.Println(a)
			case <-done:          // читаем сигнал к завершению
				result <- b
				return
			}
		}
	}()
	return result
}
//_______________________________________________________________________________________


/*
Необходимо написать функцию func merge2Channels(fn func(int) int, in1 <-chan int, in2 <- chan int, out chan<- int, n int).

Описание ее работы:
n раз сделать следующее
прочитать по одному числу из каждого из двух каналов in1 и in2, назовем их x1 и x2.
вычислить f(x1) + f(x2)
записать полученное значение в out
Функция merge2Channels должна быть неблокирующей, сразу возвращая управление.

Функция fn может работать долгое время, ожидая чего-либо или производя вычисления.
Формат ввода:
количество итераций передается через аргумент n.
целые числа подаются через аргументы-каналы in1 и in2.
функция для обработки чисел перед сложением передается через аргумент fn.
Формат вывода:
канал для вывода результатов передается через аргумент out.
*/
func merge2Channels(fn func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {
	
	ch1 := make(chan chan int, n)
	ch2 := make(chan chan int, n)
	

	prom := func(ch <-chan int, in chan chan int) {
		for i := 0; i < n; i++ {
			prch := make(chan int)
			in <- prch
			go func(c chan int, a int) {
				c <- fn(a)
			}(prch, <-ch)
		}
	}

	go prom(in1, ch1)
	go prom(in2, ch2)

	go func() {
	
		for i := 0; i < n; i++ {
			out <- (<-<-ch1) + (<-<-ch2)

		}

	}()

}
const N = 20

func mainGorutine1() {

	fn := func(x int) int {
		// time.Sleep(time.Duration(rand.Int31n(N)) * time.Second)
		return x
	}
	in1 := make(chan int, N)
	in2 := make(chan int, N)
	out := make(chan int, N)

	start := time.Now()
	merge2Channels(fn, in1, in2, out, N)
	for i := 0; i < N+1; i++ {
		in1 <- i
		in2 <- i
	}

	orderFail := false
	EvenFail := false
	for i, prev := 0, 0; i < N; i++ {
		c := <-out
		if c%2 != 0 {
			EvenFail = true
		}
		if prev >= c && i != 0 {
			orderFail = true
		}
		prev = c
		fmt.Println(c)
	}
	if orderFail {
		fmt.Println("порядок нарушен")
	}
	if EvenFail {
		fmt.Println("Есть не четные")
	}
	duration := time.Since(start)
	if duration.Seconds() > N {
		fmt.Println("Время превышено")
	}
	fmt.Println("Время выполнения: ", duration)
}
//_____________________________________________________________________________________