package algorithm

func containsNearbyDuplicate(nums []int, k int) bool {
	data := make(map[int]int,0)
	for index, num := range nums {
		idx, ok := data[num]
		if !ok {
			data[num] = index
			continue
		}
		if Abs(idx-index) <= k {
			return true
		}
		data[num] = index
	}
	return false
}
