package stepik


import (
	"fmt"
	//"strings"
	"time"
)
/*
На стандартный ввод подается строковое представление даты и времени в следующем формате:

1986-04-16T05:20:00+06:00
Ваша задача конвертировать эту строку в Time, а затем вывести в формате UnixDate:

Wed Apr 16 05:20:00 +0600 1986
*/

func OkOk(){

t, err := time.Parse(time.RFC3339 , InputTime()) 
if err != nil {
	fmt.Println(err)
}
res := t.Format(time.UnixDate)
//fmt.Println(t)
fmt.Println(res)

}


func InputTime () string{
	return "1986-04-16T05:20:00+06:00"
}
