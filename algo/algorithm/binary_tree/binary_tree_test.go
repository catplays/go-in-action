package binary_tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVerifySeriz(t *testing.T) {
	ass := assert.New(t)
	ass.Equal(true, isValidSerialization("9,3,4,#,#,1,#,#,2,#,6,#,#"))
	ass.Equal(false, isValidSerialization("1,#"))
	ass.Equal(false, isValidSerialization("9,#,#,1"))
}
