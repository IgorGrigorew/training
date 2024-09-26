package lessonyandex

import (
	"fmt"
	
	
)

/* Создайте структуру Person с полями name, age и address.
 Создайте метод Print для этой структуры, который будет
  выводить информацию о человеке на экран */
type Person struct{
	name string; 		
	age int;			
	address string;	
}

//Mehtod Print
func (p Person) Print(){
	fmt.Println("Name:", p.name)
	fmt.Println("Age:", p.age)
	fmt.Println("Address:", p.address)
}


//__________________________________________________
/* Создайте структуру Employee с полями name (string), 
position (string), salary (int) и bonus (int). 
*/
type Employee struct{
	name string;
	position string;
	salary float64;
	bonus float64;
}

/*
Создайте метод CalculateTotalSalary для этой структуры,
 который будет возвращать общую зарплату работника (salary + bonus). 
 */
 func (e Employee) CalculateTotalSalary() {
	fmt.Println("Employee: ",e.name)
	fmt.Println("Position: ",e.position)
	fmt.Printf("Total Salary: %.2f",e.salary + e.bonus)
		
	}
//_____________________________________________________________________________________________
/*
 Создайте интерфейс Shape с методом Area, который будет возвращать площадь фигуры.
*/
type Shape interface {
	Area() float64
}

 /*
 Создайте структуры Circle (с полем radius float64) и Rectangle (с полями width и height float64),
  которые будут реализовывать этот интерфейс. */
type Circle struct{
	radius float64
}

func (c Circle) Area() float64{
	const Pi float64 = 3.14159265358979323846264
	return Pi * (c.radius * c.radius)
}

type Rectangle struct{
	width float64;
	height float64
}

func (r Rectangle) Area() float64{
	return r.width * r.height
}
/*
func main() {
	r := Rectangle{width: 5.0, height:  4.0}
c := Circle{radius: 6.0}
	var s Shape  

	fmt.Println(s.Area())
	s = r
	fmt.Println(s.Area())

}
*/
//________________________________________________________________________________________


/* Создайте интерфейс Animal с методом MakeSound, который будет выводить звук,
 издаваемый животным. Создайте структуры Dog и Cat,
 которые будут реализовывать этот интерфейс и издавать соответствующие звуки
  (выводить на экран) – "Гав!" и "Мяу!". */
type Animal interface{
	MakeSound()
}

type Dog struct{}

func (d Dog) MakeSound(){
	fmt.Println("Гав!")
}

type Cat struct{}

func (c Cat) MakeSound(){
	fmt.Println("Мяу!")
}

/*
func main() {
	
	var a Animal 
	var b Animal
	a = Dog{}
	b = Cat{}
	a.MakeSound()
	b.MakeSound()
}
*/
//______________________________________________________________________________

type Student struct{
	//name string;
	solvedProblems int;    //количество решённых задач
	scoreForOneTask float64;    //количество баллов за одну задачу
	passingScore float64 		//проходной балл в следующий модуль
}

func (s Student)IsExcellentStudent() bool{

	return s.scoreForOneTask * float64(s.solvedProblems) > s.passingScore
} 

//__________________________________________________________________________
/*
Создайте интерфейс Logger с методом Log, который будет записывать сообщение в журнал.
 Создайте пользовательский тип LogLevel типа string,
 сделайте константные значения типа LogLevel со значениями Error и Info. Создайте структуру Log с полем LogLevel.
  Реализуйте метод Log c параметром типа string (текст ошибки),
 который в зависимости от текущего LogLevel выводит сообщение "ERROR: *текст ошибки*" или "INFO: *текст ошибки*".
*/

type Logger interface{
	Log()
}

type LogLevel string

const Error LogLevel = "ERROR:"

const Info LogLevel = "INFO:"

type Log struct{
	Level LogLevel
}

func (l Log)Log(message string){
fmt.Println(l.Level, message)
}