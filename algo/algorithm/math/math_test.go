package math

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidUtf8(t *testing.T) {
	ass:= assert.New(t)
	fmt.Println(1<<7)
	ass.Equal(false, validUtf8([]int{250,145,145,145,145}))
	ass.Equal(false, validUtf8([]int{145}))
	ass.Equal(true, validUtf8([]int{197,130,1}))
	ass.Equal(false, validUtf8([]int{235,140,4}))

}
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

func TestSingleNon(t *testing.T) {
	arr := []int{1,1,2,3,3,4,4,8,8}
	ass := assert.New(t)
	ass.Equal(2, singleNonDuplicate(arr))
}

func TestContainsNearbyAlmostDuplicate(t *testing.T) {
	arr := []int{2147483640,2147483641}
	ass := assert.New(t)
	ass.Equal(true, containsNearbyAlmostDuplicate(arr, 1, 100))
}


func TestThirdMax(t *testing.T) {
	arr := []int{2,2,3,1}
	ass := assert.New(t)
	ass.Equal(1, thirdMax(arr))

	arr = []int{2,2,3,1}
	ass.Equal(1, thirdMax(arr))
	arr = []int{2,3}
	ass.Equal(3, thirdMax(arr))

}