package main

import (
	"fmt"
	"strconv"

	
	"net/http"
)

var count int

func main() {
	http.HandleFunc("/count", handlereg)

	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		fmt.Println(err)
	}
}

func handlereg(w http.ResponseWriter, r *http.Request) {
	
	switch r.Method {
	case "GET":
		
		w.Write([]byte(strconv.Itoa(count)))

	case "POST":
		
		 r.ParseForm()
		 
		str := r.FormValue("count")
		num, err := strconv.Atoi(str)
		if err != nil {
			w.Write([]byte("это не число"))
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	count += num
		w.Write([]byte(strconv.Itoa(count)))

	default:
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
	}
	
	}

