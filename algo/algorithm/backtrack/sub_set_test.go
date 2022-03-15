package backtrack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountMaxOrSubsets(t *testing.T) {
	ass := assert.New(t)
	ass.Equal(2, countMaxOrSubsets([]int{3,1}))
	ass.Equal(7, countMaxOrSubsets([]int{2,2,2}))
}
