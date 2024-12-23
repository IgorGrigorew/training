package main


import (
	//"fmt"
	"time"
)



// gather выполняет переданные функции одновременно
// и возвращает срез с результатами, когда они готовы
func gather(funcs []func() any) []any {
	// начало решения

	
	results := make([]any, len(funcs))

	
    ch := make(chan struct {
        index int
        result any
    }, len(funcs))

    for i, f := range funcs {
        go func(i int, f func() any) {
            ch <- struct {
                index int
                result any
            }{index: i, result: f()}
        }(i, f)
    }

    for range funcs {
        res := <-ch
        results[res.index] = res.result

	}





	// выполните все переданные функции,
	// соберите результаты в срез
	// и верните его
	

return results
	// конец решения
}

// squared возвращает функцию,
// которая считает квадрат n
func squared(n int) func() any {
	return func() any {
		time.Sleep(time.Duration(n) * 100 * time.Millisecond)
		return n * n
	}
}
/*
func main() {
	funcs := []func() any{squared(2), squared(3), squared(4)}

	start := time.Now()
	nums := gather(funcs)
	elapsed := float64(time.Since(start)) / 1_000_000

	fmt.Println(nums)
	fmt.Printf("Took %.0f ms\n", elapsed)
}
*/
