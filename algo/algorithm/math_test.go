package algorithm

import (
	"fmt"
	"testing"
)

func TestCanFinish(t *testing.T) {
	a := [][]int {{1,0},{0,1}}
	var b = CanFinish(2, a)
	fmt.Println(b)
}
func TestIsIsomorphic(t *testing.T)  {
	b := IsIsomorphic("bbbaaaba","aaabbbba")
	fmt.Println(b)
}

func TestSummaryRanges(t *testing.T)  {
	fmt.Println(summaryRanges([]int{0,1,2,4,5,7}))
	fmt.Println(summaryRanges([]int{0,2,3,4,6,8,9}))
}