package main

import (
	"bufio"
	"slices"
	"time"

	"fmt"
	"os"
	"strconv"
)
// закрытые тесты не проходит!!!
type In struct {
	A int
	B int
	C int
	N int
}

func main() {
	start := time.Now()
	fmt.Println("Start")
	yandex()
	end := time.Now()

	fmt.Println("Stop", end.Sub(start))
}

func yandex() {

	intstr := input()

	n := intstr.N
	
	var min, avg, max int
	sl := make([]int, 3)
	sl[0] = intstr.A
	sl[1] = intstr.B
	sl[2] = intstr.C

	slices.Sort(sl)
	min, avg, max = sl[0], sl[1], sl[2]

	chn := make(chan []int)

	m := make(map[int]bool)
	i := min

	for i < n {

		go sravn(min, avg, max, i, n, chn)

		i += min
for _, v := range <- chn {
_, ok := m[v]
if !ok {
	m[v] = true
//	fmt.Println(v)
}


}
		
//break
	}

	fmt.Println("output",len(m))

}

func sravn(a, b, c, i, n int, ch chan []int) {

	/*
res  := make([]int, 3)
resab := make([]int, n/a)
resac := make([]int, n/a)
resabc := make([]int, n/a)
resacb := make([]int, n/a)



if i < n {
res[0] = i
}
if b * (i/a) < n {
	res[1] = b * (i/a)
	}
if c * (i/a) < n {
		res[2] = c * (i/a)
		}

for y:= 1 ;  ; y++ {
	

ab := i + (b * y)
if ab < n {
resab[y-1] = ab
	
}else{
	break
}
ac := i + (c * y)
if ac < n {
	resac[y-1] = ac
}
		for j := 1; ; j ++ {
		ab += c * j
		if ab < n {
		resabc[j-1] = ab
		}
		ac += b * j
		if ac < n {
			resacb[j-1] = ac
			}else{
				break
			}
		}

}

res = append(res, resab...)
res = append(res, resac...)
res = append(res, resabc...)
res = append(res, resacb...)
*/

res := make([]int, 18)
del := i/a

if i < n {
	res[0] = i
}
if i + b < n {
	res[1] = i + b
}
if i + c < n {
	res[2] = i + c
}



if b*del < n {
	res[3] = b*del
}
if (b*del) + a  < n {
	res[4] = (b*del) + a
}
if (b*del) + c < n {
	res[5] = (b*del) + c
}
if (b*del) + a + c  < n {
	res[6] = (b*del) + a + c
}
if (b*del) + i  < n {
	res[7] = (b*del) + i
}
if (b*del) + i + c  < n {
	res[8] = (b*del) + i + c
}
if (b*del) + i + (c*del)  < n {
	res[9] = (b*del) + i + (c*del)
}
if (b*del)  + (c*del)  < n {
	res[10] = (b*del)  + (c*del)
}
if (b*del)  + (c*del) + a  < n {
	res[11] = (b*del)  + (c*del) + a
}


if c*del < n {
	res[12] = c*del
}
if (c*del) + a < n {
	res[13] = (c*del) + a
}
if (c*del) + b < n {
	res[14] = (c*del) + b
}
if (c*del) + a + b < n {
	res[15] = (c*del) + a + b
}
if (c*del) + i < n {
	res[16] = (c*del) + i
}
if (c*del) + i + b < n {
	res[17] = (c*del) + i + b
}




ch <- res
	
}

func input() In {

	file, err := os.Open("text.txt")
	if err != nil {
		fmt.Print(err)
	}

	var in *bufio.Reader
	var out *bufio.Writer
	//	in = bufio.NewReader(os.Stdin)

	in = bufio.NewReader(file)

	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, a, b, c string
	var inIn In

	fmt.Fscan(in, &n)
	fmt.Fscan(in, &a, &b, &c)
	numN, _ := strconv.Atoi(n)
	numA, _ := strconv.Atoi(a)
	numB, _ := strconv.Atoi(b)
	numC, _ := strconv.Atoi(c)

	inIn.A = numA
	inIn.B = numB
	inIn.C = numC
	inIn.N = numN

	//	fmt.Fprint(out, inIn)

	return inIn
}
