package lessonyandex

import (
	"fmt"
	"time"
)


/*
func main() {
e := Employee{
	name:"Anton",
	position: "product manager", 
	salary: 100,
	bonus: 100,

}
	e.CalculateTotalSalary()

}
*/


type Note struct{
	title string; 		//заголовок 
	text string 		//текст заметок
}

type ToDoList struct {
	name string; 		//название списка
	tasks []Task; 		//список дел на сегодня
	notes []Note; 		//список дополнительных заметок

}

// общее количество задач
func (t ToDoList)TasksCount() int{

return len(t.tasks)
}

// общее количество заметок 
func (t ToDoList)NotesCount() int{

	return len(t.notes)

}

// количество приоритетных задач
func (t ToDoList)CountTopPrioritiesTasks() int{
res := 0

for _, tas := range t.tasks {

	if tas.IsTopPriority(){
		res ++
	}
}

return res
}


// количество просроченных задач
func (t ToDoList)CountOverdueTasks() int {
ovT := 0

for _, ta := range t.tasks {

	if !ta.IsOverdue() {
		ovT ++
	}

}

return ovT
}



//_______________________________________________

type Task struct{
	summary string;  //короткий заголовок
	description  string ; //подробное описание
	deadline time.Time;   //дедлайн для задачи 
	priority int 			//приоритет задачи
}

func (t Task)IsOverdue()bool{

timeN := time.Now()

return timeN.After(t.deadline)
}


func (t Task)IsTopPriority()bool{

	return t.priority > 3
}




//__________________________________________
type Employee1 struct{
	name string;
	position string;    
	salary float64;
	bonus float64;
}

/*
Создайте метод CalculateTotalSalary для этой структуры,
 который будет возвращать общую зарплату работника (salary + bonus). 
 */
func (e Employee) CalculateTotalSalary1() {
fmt.Println("Employee: ",e.name)
fmt.Println("Position: ",e.position)
fmt.Printf("Total Salary: %.2f",e.salary + e.bonus)
	
}
