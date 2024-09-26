package lessoncodebasic

//Реализуйте функции nextASCII(b byte) byte и prevASCII(b byte) byte,
// которые возвращают следующий или предыдущий символ ASCII таблицы соответственно

//возвращают следующий символ ASCII таблицы
func NextASCII(b byte) byte {
	return b + 1
}

//возвращают предыдущий символ ASCII таблицы
func PrevASCII(b byte) byte {
	return b - 1
}
