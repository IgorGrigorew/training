package stepik

import (
	"fmt"
	"strings"
	//"strings"
	"time"
)

/*
На стандартный ввод подается строковое представление двух дат,
	разделенных запятой (формат данных смотрите в примере).
Необходимо преобразовать полученные данные в тип Time, а затем вывести продолжительность
	периода между меньшей и большей датами.
*/
func TimeDateOk() {
	str := strings.Split(InputTimeDate(), ",")
	
layot := "2.01.2006 15:04:05"

	t1, err := time.Parse( layot, str[0])
	if err != nil {
		fmt.Println(err)
	}
	t2, err := time.Parse(layot, str[1])
	if err != nil {
		fmt.Println(err)
	}
	if t1.Before(t2){
	fmt.Println(t2.Sub(t1))
	}else{
		fmt.Println(t1.Sub(t2))
	}
}

func InputTimeDate() string {
	return "11.03.2018 14:00:15,12.03.2018 14:00:15"
}
