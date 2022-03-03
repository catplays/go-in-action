package dp

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinSubArrayLen(t *testing.T) {
	ass := assert.New(t)
	//ass.Equal(45,computeArea(-3,0,3,4,0,-1,9,2))
	ass.Equal(17,computeArea(-2,-2,2,2,-3,-3,-2,-2))

}
