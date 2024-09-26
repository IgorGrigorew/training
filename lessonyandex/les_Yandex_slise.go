package lessonyandex

import (
	"fmt"
)


func UnderLimit(nums []int, limit int, n int) ([]int, error) {
	res := []int{}

	if n < 0 || nums == nil {
		return nil, fmt.Errorf("error")
	}

	for _, v := range nums {

		if v < limit {
			res = append(res, v)
		}

	}

	if len(res) > n {
		return res[:n], nil
	}

	return res, nil
}

//________________________________________________________

func Clean(nums []int, x int) []int {

	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] == x {
			nums = append(nums[:i], nums[i+1:]...)

		}

	}
	return nums
}

//_____________________________________________________

func SliceCopy(nums []int) []int {
	var res = make([]int, len(nums))

	for i, v := range nums {
		res[i] = v
	}
	return res
}

//_________________________________________________________-

func Join(nums1, nums2 []int) []int{
	c := cap(nums1) + cap(nums2)
res := make([]int, 0, c)
res = append(res, nums1...)

res = append(res, nums2...)

return res

}

//__________________________________________________

func  Mix(nums []int) []int{
res := []int{}

for i, y := 0, (len(nums)/2); y < len(nums); i++ {
res = append(res, nums[i], nums[y])
y++
}

return res
}