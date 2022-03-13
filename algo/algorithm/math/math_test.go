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
