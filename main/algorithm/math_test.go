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
