package lessoncodebasic

//Реализуйте тип PersonList (слайс структур Person), 
// c методом (pl PersonList) GetAgePopularity() map[uint8]int, который возвращает мапу,
//где ключ — возраст, а значение — кол-во таких возрастов

type Person struct {
    Age uint8
}

type PersonList []Person

func (pl PersonList) GetAgePopularity() map[uint8]int{
	mapa:= make(map[uint8]int)
	for _, i := range pl {
		if _, ok := mapa[i.Age]; ok {
mapa[uint8(i.Age)]++
		}else{
mapa[uint8(i.Age)]=1
		} 
	}
	return mapa
}