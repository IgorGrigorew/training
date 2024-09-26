package lessoncodebasic

import "sort"

/* Реализуйте функцию UniqueSortedUserIDs(userIDs []int64) []int64, которая возвращает отсортированный слайс,
состоящий из уникальных идентификаторов userIDs.
 Обработка должна происходить in-place, то есть без выделения доп. памяти. */

func UniqueSortedUserIDs(userIDs []int64) []int64 {
	sort.Slice(userIDs, func(i, j int) bool { return userIDs[i] < userIDs[j] })
	resuserId := make([]int64, 0)
	if len(userIDs) == 0 {
		return resuserId
	}
	resuserId = append(resuserId, userIDs[0])
	for i := 1; i < len(userIDs); i++ {
		if userIDs[i] == userIDs[i-1] {
			continue
		} else {
			resuserId = append(resuserId, userIDs[i])
		}
	}
	return resuserId
}
