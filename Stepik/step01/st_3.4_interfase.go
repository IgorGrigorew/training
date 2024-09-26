package stepik


import (
	"fmt"
	"strings"
)
/*
Давайте используем ваши знания структур, методов и интерфейсов на практике и реализуем объект,
 удовлетворяющий интерфейсу fmt.Stringer. Назовем его "Батарейка".

Во-первых, вы должны объявить новый тип, удовлетворяющий интерфейсу fmt.Stringer.

Ваш тип должен предусматривать, что на печати он будет выглядеть так: [      XXXX]:
 где пробелы - "опустошенная" емкость батареи, а X - "заряженная".

Во-вторых, на стандартный ввод вы получаете строку, состоящую ровно из 10 цифр: 0 или 1 (порядок 0/1 случайный).
 Ваша задача считать эту строку любым возможным способом и создать на основе этой строки 
 объект объявленного вами на первом этапе типа: надеюсь, вы понимаете,
  что строка символизирует емкость батарейки: 0 - это "опустошенная" часть, а 1 - "заряженная".

В-третьих, созданный вами объект должен называться batteryForTest (использование этого имени обязательно).

В вашем распоряжении фактически весь файл, НО завершающая фигурная скобка функции main() вам не видна, 
но она присутствует. Перед этой скобкой присутствует функция (которая вам тоже не видна),
 принимающая в качестве аргумента объект типа fmt.Stringer - batteryForTest,
  и направляющая его на стандартный вывод, поэтому вам не требуется выводить что-то на печать самостоятельно.
*/
type Battery struct{
	Charge  string 
}

func (b *Battery) String () string {
str:= b.Charge
lenStr := len(str)
result := make([]string, lenStr)
var left int
rigt := lenStr-1

for i:= 0; left<=rigt; i++ {
if string(str[i]) == "0" {
result[left]=" "
left++
}else{
	result[rigt]="X"
	rigt--
}
}

	return "[" + strings.Join(result, "") + "]"
}



func Ok() {
inp:= Inpun()

b := &Battery{Charge: inp}
batteryForTest := b.String()
fmt.Println(batteryForTest)

}


func Inpun () string{
	return "0000010001"
}
