package main

import (

	
	"fmt"
	"log"
	"net"
	"strings"
	//"time"
)

/*
Подключитесь к адресу 127.0.0.1:8081 по протоколу TCP, считайте от сервера 3 сообщения,
 и выведите их в верхнем регистре. Рекомендуем использовать буфер в 1024 байта.
*/




func main() {
	//("udp", nil, &net.UDPAddr{IP: []byte{127, 0, 0, 1}, Port: 8081})
	conn, err := net.DialTCP("tcp", nil, &net.TCPAddr{IP: []byte{127,0,0,1}, Port: 8081})
	if err != nil {
		log.Fatalln(err)
	}
	//fmt.Println("Start client")
	
	go func (){
for {
	buf := make([]byte, 1024)
n, err := conn.Read(buf)
if err != nil {
	log.Println(err)
	break
}
fmt.Println("Получено: ", strings.ToTitle(string(buf[:n])))
}

//fmt.Println("Stop client")
defer conn.Close()
	}()
	//time.Sleep(time.Second*3)

}

//_________________________________________________________________________________
//server//

/*
package main

import (
	//"container/list"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
	//"time"
)

func main() {
	//fmt.Println("Start")
// 127.0.0.1:8081
	//("udp", &net.UDPAddr{IP: []byte{127, 0, 0, 1}, Port: 8081})
		serverConn, err := net.ListenTCP("tcp", &net.TCPAddr{IP: []byte{127,0,0,1}, Port: 8081})
		if err != nil {
			log.Fatalln(err)
		}
		//fmt.Println("Conecting")
		defer serverConn.Close()
		con, err := serverConn.AcceptTCP()
			if err != nil {
				log.Println(err)
			}

con.Write([]byte("messeg1"))
time.Sleep(10)

con.Write([]byte("messeg2"))
time.Sleep(10)

con.Write([]byte("messeg3"))


		go func(){
		 ServerConn(con)
	}()

time.Sleep(time.Second*5)
fmt.Println("End")

}


func ServerConn(con *net.TCPConn) {
//	fmt.Println("Обработка сообщения")
	for {
		buf := make([]byte, 1024)
		n, err := con.Read(buf)
		if err != nil {
			log.Print(err)
			break
		}


		fmt.Println("Получено: ", strings.ToTitle(string(buf[:n])))
	}
	defer func() {
		if err := con.Close(); err != nil {
		 log.Printf("Error closing connection: %v", err)
		}
	   }()	
}
*/