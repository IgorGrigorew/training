package lessonyandex

/*Напишите функцию CalculateDigitalRoot(n int) int,
 которая будет принимать на вход положительное целое
  число n и возвращать его цифровой корень. */
  func CalculateDigitalRoot(n int) int {
	res := 0
	res = Sum(n)
	for {
	if res > 9 {
		res = Sum(res)
		continue
	}else{
		break
	}
	}
	return res
	}
	
	func Sum(s int) int{
		res := 0
		for i:= s; i > 0 ; i /= 10 {
			res += i%10
				}
		return res
	}
	