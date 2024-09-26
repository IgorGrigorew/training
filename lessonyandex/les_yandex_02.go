package lessonyandex

import "fmt"

//выводит на экран слово YES, если число N является точной степенью двойки, или слово NO в противном случае
func IsPowerOfTwoRecursive(N int) {
	var a int = 2
	var res int = 1
	for i := 1; ; i++ {
		res *= a
		if N == 1 {
			fmt.Print("YES")
			break
		} else	if res > N {
			fmt.Print("NO")
			break
		} else if res == N {
			fmt.Print("YES")
			break
		}
	}
}
