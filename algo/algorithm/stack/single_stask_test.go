package stack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxArea(t *testing.T) {
	arr := [][]byte{{1,0,1,0,0},{1,0,1,1,1},{1,1,1,1,1},{1,0,0,1,0}}
	ass := assert.New(t)
	ass.Equal(6, maximalRectangle(arr))
}
