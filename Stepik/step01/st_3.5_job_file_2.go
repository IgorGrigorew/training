package stepik

import (
	//"bytes"
	"encoding/csv"
	"fmt"
	//"io/ioutil"
	"os"

	"path/filepath"
	"strings"
)
/*
Данная задача поможет вам разобраться в пакете encoding/csv и path/filepath,
 хотя для решения может быть использован также пакет archive/zip 
 (поскольку файл с заданием предоставляется именно в этом формате).
В тестовом архиве, который вы можете скачать из нашего репозитория на github.com, 
содержится набор папок и файлов. Один из этих файлов является файлом с данными в формате CSV, 
прочие же файлы структурированных данных не содержат.
Требуется найти и прочитать этот единственный файл со структурированными данными
(это таблица 10х10, разделителем является запятая), а в качестве ответа необходимо указать число, 
находящееся на 5 строке и 3 позиции (индексы 4 и 2 соответственно).
*/
func mywalkFunc(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	if strings.Contains(info.Name(), "file") {
		
file, err := os.Open(path)
		if err != nil {
			panic(err)
		}
defer file.Close()
	csvBuffer := csv.NewReader(file)
	strSl, _ := csvBuffer.ReadAll()

if len(strSl)>1 {
	fmt.Println(strSl[4][2])
}

		return nil
	} else {
		return nil
	}
}


/*
func main() {
	*/
	func Start (){
	const root = "./Stepik/task"
	
	if err := filepath.Walk(root, mywalkFunc); err != nil {
		fmt.Printf("ошибка: %v ", err)
	}
	
}

