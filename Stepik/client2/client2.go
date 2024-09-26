package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)
/*
Считайте с консоли две переменные, сначала name, затем age. 
Сделайте HTTP запрос на http://127.0.0.1:8080/hello передав name и age в query параметрах,
 затем напечатайте ответ сервера (только тело).
*/

func main() {
//	adres := "stepik.org"
	adres := "http://127.0.0.1:8080/hello"
	parametrs := url.Values{}
	parametrs.Add("age", "20")
	parametrs.Add("name", "Dima")
	
	fulladres := adres + "?" + parametrs.Encode()

// отправляем запрос
resp, err := http.Get(fulladres)
	if err != nil {
		log.Fatalln(err)
	}
defer resp.Body.Close()

byteresp, err := io.ReadAll(resp.Body)
if err != nil {
	log.Fatalln(err)
}
fmt.Println(string(byteresp))

}