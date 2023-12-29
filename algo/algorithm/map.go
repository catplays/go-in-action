package algorithm

import "catwang.com/go-in-action/algo/algorithm/math"

func containsNearbyDuplicate(nums []int, k int) bool {
	data := make(map[int]int,0)
	for index, num := range nums {
		idx, ok := data[num]
		if !ok {
			data[num] = index
			continue
		}
		if math.Abs(idx-index) <= k {
			return true
		}
		data[num] = index
	}
	return false
}
