package algorithm

import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
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

func TestMinSubArrayLen(t *testing.T) {
	ass := assert.New(t)
	ass.Equal(2,minSubArrayLen(7,[]int{2,3,1,2,4,3}))

	ass.Equal(3,minSubArrayLen(11,[]int{1,2,3,4,5}))

	ass.Equal(0,minSubArrayLen(11,[]int{1,1,1,1,1,1,1,1}))
}