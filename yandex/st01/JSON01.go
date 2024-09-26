package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func main() {



m, err := modifyJSON(input())
if err != nil{
	fmt.Println(err)
}
fmt.Println(string(m))
}

type Student struct {
    Name  string `json:"name"`
    Grade int    `json:"grade"`
}



func input () []byte{
j, err := os.Open("text.txt")
	if err != nil {
		fmt.Println(err)
	}
defer j.Close()

byt, err  := io.ReadAll(j)
if err != nil{
	fmt.Println(err)
}

	return byt
}


//принимает данные в формате JSON, добавляет 1 год к полю grade 
//и возвращает обновлённые данные также в формате JSON.

func modifyJSON(jsonData []byte) ([]byte, error){
var stud []Student

err := json.Unmarshal(jsonData, &stud)
if err != nil{
	fmt.Println("error 01",err)
}
fmt.Println(stud)


for i, st := range stud  {
	
	stud[i].Grade = st.Grade + 1

	//fmt.Println(st)
}




fmt.Println(stud)


res, err := json.Marshal(stud)
if err != nil{
	fmt.Println("error 02",err)
}

return res, nil
}