package stack

import "fmt"

type IStack interface {
	Push(val int)
	Pop() (int, error)
	Top() (int, error)
	IsEmpty() bool
}

type Stack struct {
	capacity int
	top int
	data []int
}

func NewStack(capacity int) Stack {
	return Stack{
		capacity: capacity,
		top: 0,
		data: make([]int, 0),
	}
}
func (s *Stack) Push(val int) {
	s.data = append(s.data, val)
	s.top++
}

func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, fmt.Errorf("stask empty")
	}
	val := s.data[s.top-1]
	s.data = s.data[:s.top-1]
	s.top--
	return val, nil
}

func (s Stack) Top() (int, error) {
	if s.IsEmpty() {
		return 0, fmt.Errorf("stask empty")
	}
	return s.data[s.top-1], nil
}

func (s Stack) IsEmpty() bool {
	return s.top == 0
}

