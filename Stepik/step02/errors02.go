package main

// не меняйте импорты, они нужны для проверки
import (
   // "bufio"
  //  "errors"
    "fmt"
  //  "io/ioutil"
  //  "os"
 //   "reflect"
  //  "runtime"
    "strconv"
    "strings"
)

// account представляет счет
type account1 struct {
    balance   int
    overdraft int
}

var myerr error

func mainErr() {
    var acc account1
    var trans []int
	
    acc, trans = parseInput1()
        fmt.Print("-> ")
      err := myerr  
        if err != nil {
            fmt.Println(err)
            return
        }
        fmt.Println(acc, trans)
   
	

    
	
}

// parseInput считывает счет и список транзакций из os.Stdin.
func parseInput1() (account1, []int) {
    accSrc, transSrc := readInput1()
    acc := parseAccount(accSrc)
    trans := parseTransactions(transSrc)
    return acc, trans
}

// readInput возвращает строку, которая описывает счет
// и срез строк, который описывает список транзакций.
// эту функцию можно не менять
func readInput1() (string, []string) {
	/*
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Split(bufio.ScanWords)
    scanner.Scan()
	
*/
	//80/x 10 y 30
    accSrc := "80/10"
      transSrc := []string{"10", "y", "30"}
	 /*
    for scanner.Scan() {
        transSrc = append(transSrc, scanner.Text())
    }
	*/
    return accSrc, transSrc
}

// parseAccount парсит счет из строки
// в формате balance/overdraft.
func parseAccount(src string) account1 {
    parts := strings.Split(src, "/")
    balance, err := strconv.Atoi(parts[0])
		if err != nil{
			myerr = err
			return account1{}
		}
    overdraft, err := strconv.Atoi(parts[1])
	if err != nil{
		myerr = err
		return account1{}
	}
    if overdraft < 0 {
        myerr = fmt.Errorf("expect overdraft >= 0")
    }
    if balance < -overdraft {
		myerr = fmt.Errorf("balance cannot exceed overdraft")
    }
    return account1{balance, overdraft}
}

// parseTransactions парсит список транзакций из строки
// в формате [t1 t2 t3 ... tn].
func parseTransactions(src []string) []int {
    trans := make([]int, len(src))
    for idx, s := range src {
        t, err := strconv.Atoi(s)
		if err != nil && myerr == nil {
			myerr = err
		}
        trans[idx] = t
    }
    return trans
}