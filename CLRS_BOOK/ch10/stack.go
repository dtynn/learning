package ch10

import "fmt"

func NewStack(capacity int) *Stack {
	return &Stack{
		stack: make([]int, capacity),
	}
}

type Stack struct {
	stack []int
	head  int
}

func (s *Stack) Push(x int) error {
	if s.head >= len(s.stack) {
		return fmt.Errorf("stack overflow")
	}

	s.stack[s.head] = x
	s.head++
	return nil
}

func (s *Stack) Pop() (int, error) {
	if s.head == 0 {
		return NIL, fmt.Errorf("stack underflow")
	}

	s.head--
	x := s.stack[s.head]
	return x, nil
}
