package lessoncodebasic

// принимает на вход состоящую из ASCII символов строку s и возвращает новую строку,
// где каждый символ из входящей строки сдвинут вперед на число step

func ShiftASCII(s string, step int) string {
	bys := []byte(s)
	if step > 284 || step < -512 {
		step = 256
	}
	for i := 0; i < len(s); i++ {
		bys[i] += byte(step)
	}
	//fmt.Print(bys)
	return string(bys)
}
