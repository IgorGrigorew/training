package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {

	Input()
}

func Input() {
	/*
		var in *bufio.Reader
		var out *bufio.Writer
		in = bufio.NewReader(os.Stdin)
		out = bufio.NewWriter(os.Stdout)
		defer out.Flush()

		var a, b int
		fmt.Fscan(in, &a, &b)
		fmt.Fprint(out, a + b)
	*/

	file, err := os.Open("text.txt")
	if err != nil {
		fmt.Println(err)
	}

	var in *bufio.Reader
	var out *bufio.Writer
	/*
		in = bufio.NewReader(os.Stdin)

	*/

	in = bufio.NewReader(file)

	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var a, b string
	fmt.Fscan(in, &a)

	num, err := strconv.Atoi(a)
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	for i := 0; i < num; i++ {

		fmt.Fscan(in, &a, &b)

		n, err := strconv.Atoi(a)
		if err != nil {
			fmt.Println(err)
		}
		prc, err := strconv.Atoi(b)
		if err != nil {
			fmt.Println(err)
		}
		slfloat := make([]float64, n)
		for y := 0; y < n; y++ {
			fmt.Fscan(in, &a)
			nfloat, err := strconv.ParseFloat(a, 64)
			if err != nil {
				fmt.Println(err)
			}
			slfloat[y] = nfloat
		}

		res := Prcent(prc, slfloat...)

		fmt.Fprintf(out, "%.2f\n", res)
	}

}

func Prcent(prc int, prise ...float64) float64 {
	var res, prom float64
	for _, pri := range prise {
		prom = pri * float64(prc) / 100
		if prom < 1.00 {
			res += prom
		} else {
			res += prom - math.Floor(prom)
		}

	}

	return res
}
