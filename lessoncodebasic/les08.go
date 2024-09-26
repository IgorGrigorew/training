package lessoncodebasic

//  возвращает true,  если строка s состоит только из ASCII символов
func IsASCII(s string) bool {
	for _, i := range s {
		if i > 127 {
			return false
		}
	}
	return true
}
