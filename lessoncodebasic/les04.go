package lessoncodebasic

//возвращает слайс, состоящий из уникальных идентификаторов userIDs.
//Порядок слайса должен сохраниться.
func UniqueUserIDs(userIDs []int64) []int64 {
	resuserId := make([]int64, 0)
	if len(userIDs) == 0 {
		return resuserId
	}
	resuserId = append(resuserId, userIDs[0])
	for i := 1; i < len(userIDs); i++ {
		flag := false
		for y := len(resuserId); y > 0; y-- {
			if userIDs[i] == resuserId[y-1] {
				flag = false
				break
			}
			flag = true
		}
		if flag {
			resuserId = append(resuserId, userIDs[i])
		}
	}

	return resuserId
}
