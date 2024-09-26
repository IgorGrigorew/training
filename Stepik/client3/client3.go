package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func main001() {
	adres := "http://127.0.0.1:3333/count"	
	parametrs := url.Values{}
	parametrs.Add("count", "go20")
	res, err := http.PostForm(adres, parametrs)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

byteresp, err := io.ReadAll(res.Body)
if err != nil {
	fmt.Println(err)
}
fmt.Println(string(byteresp))
fmt.Println("end")

}