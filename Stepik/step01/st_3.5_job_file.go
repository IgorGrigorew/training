package stepik

import (
	"encoding/csv"
	"fmt"
	"os"
)

func FileJob() {
	// чтение из файла
	f1, err := os.Open("task.data")
	if err != nil {
		fmt.Println(err)
	}
	defer f1.Close()
	//  fmt.Println(1)
	csvBuffer := csv.NewReader(f1)
	csvBuffer.Comma = ';'
	record, _ := csvBuffer.Read() // Reads 1 line

	//fmt.Println(record)
	for col, content := range record {

		if content == "0" {
			fmt.Println(col + 1)
			break
		}
	}
}