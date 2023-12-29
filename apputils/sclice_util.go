package apputils

func ContainsInt(nums []int, i int) bool {
	for _, item := range nums {
		if item == i {
			return true
		}
	}
	return false
}
