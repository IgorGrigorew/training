package lessonyandex

import (
	
)




func FindMaxKey1(m map[int]int) int {
max := 0
maxKey := 0
for key, vol := range m {
	if vol > max {
		max = vol
		maxKey = key
	}
}
return maxKey
}

//___________________________________________

func SumOfValuesInMap1(m map[int]int) int{
res := 0

for _, vol := range m {
	res += vol
}
return res
}

//______________________________________________

func SwapKeysAndValues1(m map[string]string) map[string]string{
res := make(map[string]string, len(m))
for key, vol := range m {
	res[vol] = key
}
return res
}

//________________________________________________

func CountingSort(contacts []string) map[string]int{
    m := make(map[string]int)
    
    for _, v := range contacts {
        _, ok := m[v]
        if ok {
            m[v] ++
        }else{
         m[v] = 1   
        }
              
    }
    return m    
    
}


//___________________________________________________


func DeleteLongKeys(m map[string]int) map[string]int{

for key := range m {
if len(key) < 6 {
	delete(m, key)
}
}
return m
}