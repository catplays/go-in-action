package stack

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStack(t *testing.T) {
	ass := assert.New(t)
	stack := NewStack(10)
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	val, _ :=stack.Pop()
	ass.Equal(3, val)
	val, _ =stack.Pop()
	ass.Equal(2, val)
	ass.Equal(false, stack.IsEmpty())
	stack.Push(5)
	val, _ =stack.Pop()
	ass.Equal(5, val)
	val, _ =stack.Pop()
	ass.Equal(1, val)
	ass.Equal(true, stack.IsEmpty())
	val, err := stack.Pop()
	fmt.Println(err)
}
func TestMyQueue(t *testing.T) {
	ass := assert.New(t)
	queue:= Constructor()
	queue.Push(1)

	val :=queue.Pop()
	ass.Equal(1, val)

	ass.Equal(true,queue.Empty() )
}

