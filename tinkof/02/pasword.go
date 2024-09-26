package _test

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


/*
func main() {

	 tin()

}
*/
func tin()  {

str, pas, n := input()

m := make(map[byte]int, len(pas))

for _, v := range str {
_, ok := m[v]
if ok {
	m[v]++
}else{
	m[v] = 1
}
}
min := m[pas[0]]
var minSumb byte
for i, v := range m {
if v < min {
	min = v
	minSumb = i
}
}
res := []string{}
for i, v := range str {
	
if v == minSumb {
	
res = append(res, string(str[i: i+n]))
}

}

for i, v := range res {

		for _, b :=range string(pas) {
if !strings.ContainsRune(v, b){

res = append(res[:i],res[i+1:]... )
}
		}

}

var resctr string
if len(res[len(res)-1]) == n{
resctr = res[len(res)-1]
}
fmt.Println(resctr)

}




func input()([]byte, []byte, int ) {

file, _ := os.Open("input.txt")

defer file.Close()

read := bufio.NewReader(file)
/*
read := bufio.NewReader(os.Stdin)
*/
sl, _ := read.ReadBytes('\n')
sl = sl[:len(sl)-2]
sl1, _ := read.ReadBytes('\n')
sl1 = sl1[:len(sl1)-2]
n := 0
fmt.Scan(&n)
return sl, sl1, n
}