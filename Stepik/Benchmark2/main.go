package main

// не удаляйте импорты, они используются при проверке
import (
    "fmt"
  //  "math/rand"
  //  "os"
    "strings"
 //   "testing"
)

// Words работает со словами в строке.
type Words  map[string]int

// MakeWords создает новый экземпляр Words.
func MakeWords(s string) Words {
    m := make(map[string]int, len(s))
	for i, v := range strings.Fields(s){
_, ok := m[v]
if !ok {
    m[v] = i
}
    }
	 
    return m
}

// Index возвращает индекс первого вхождения слова в строке,
// или -1, если слово не найдено.
func (w Words) Index(word string) int {
  
   _, ok := w[word]
   if ok {
    return w[word]
   }
    return -1
}

func main(){
	s := "go is awesome, php is not"
	w := MakeWords(s)
	
	fmt.Println(w.Index("go"))
	// 0
	fmt.Println(w.Index("is"))
	// 1
	fmt.Println(w.Index("is not"))
	// -1
	fmt.Println(w.Index("python"))
	// -1
}