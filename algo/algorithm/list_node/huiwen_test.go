package list_node

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	head := BuildListNode([]int{1,2,2,1})
	assert := assert.New(t)
	assert.Equal(true, isPalindrome(head))
}
