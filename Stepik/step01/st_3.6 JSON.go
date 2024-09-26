package stepik

import (
	"fmt"
	"io/ioutil"
	"os"
	"encoding/json"
)

/*
Данная задача ориентирована на реальную работу с данными в формате JSON.
Для решения мы будем использовать справочник ОКВЭД (Общероссийский классификатор видов экономической деятельности),
 который был размещен на web-портале data.gov.ru.
Необходимая вам информация о структуре данных содержится в файле structure-20190514T0000.json,
 а сами данные, которые вам потребуется декодировать, содержатся в файле data-20190514T0100.json.
  Файлы размещены в нашем репозитории на github.com.
Для того, чтобы показать, что вы действительно смогли декодировать документ вам необходимо в
 качестве ответа записать сумму полей global_id всех элементов, закодированных в файле.
*/

func JSONnew (){
// чтение из файла
f1, err := os.Open("data-20190514T0100.json") 
if err != nil {
	fmt.Println(err)
}
defer f1.Close() 
//  fmt.Println(1)


n, err := ioutil.ReadAll(f1)

// чтение со стандартного ввода
//n, err := ioutil.ReadAll(os.Stdin)
if err != nil {
fmt.Println(err)
}

//	fmt.Println(n, 2)
var muStr  MuStr
if err := json.Unmarshal(n, &muStr); err!=nil{
	fmt.Println(err)
}

var res int
//fmt.Println(3)
st := muStr
for _, vol := range st{

	res += vol.Id

}

fmt.Println(res)

}




type MuStr []struct {
Id int `json:"global_id"`
}
