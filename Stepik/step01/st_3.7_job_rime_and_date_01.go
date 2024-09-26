package stepik

import (
	"fmt"
	"strings"
	"time"
    "bufio"
    "os"
)
/*
На стандартный ввод подается строковое представление даты и 
времени определенного события в следующем формате:

2020-05-15 08:00:00
Если время события до обеда (13-00), то ничего менять не нужно,
 достаточно вывести дату на стандартный вывод в том же формате.

Если же событие должно произойти после обеда, необходимо перенести его на то же время на следующий день,
 а затем вывести на стандартный вывод в том же формате.
*/

func TimeOk() {

   rd := bufio.NewReader(os.Stdin)
timeStr, err := rd.ReadString('\n')
if err != nil {
    fmt.Println(err)
}

// Удаление последнего символа при помощи пакета "strings".
timeStr = strings.TrimSuffix(timeStr, "\n")
t, err := time.Parse(time.DateTime , timeStr) 
if err != nil {
	fmt.Println(err)
}
if t.Hour()< 13 {
fmt.Println(t.Format(time.DateTime))
}else{
	fmt.Println(t.Add(time.Hour*24).Format(time.DateTime))
}


}