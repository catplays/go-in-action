package datastructure

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAllOne(t *testing.T) {
	ass := assert.New(t)
	obj := Constructor()
	obj.Inc("hello")
	obj.Inc("hello")
	obj.Inc("a")
	ass.Equal("hello", obj.GetMaxKey())
 	ass.Equal("a", obj.GetMinKey())
	obj.Dec("a")

	ass.Equal("hello", obj.GetMaxKey())
	ass.Equal("hello", obj.GetMinKey())
}
