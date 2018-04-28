package stack_with_min

type StackWithMin struct {
	vals []int
	min  []int
}

func (s *StackWithMin) Push(i int) {
	s.vals = append(s.vals, i)

	if len(s.min) == 0 {
		s.min = append(s.min, i)
		return
	}

	min := s.min[0]
	if i < min {
		min = i
	}

	s.min = append(s.min, min)
}

func (s *StackWithMin) Pop() (int, bool) {
	size := len(s.vals)
	if size == 0 {
		return 0, false
	}

	tail := s.vals[size-1]
	s.vals = s.vals[:size-1]
	s.min = s.min[:size-1]
	return tail, true
}

func (s *StackWithMin) Min() (int, bool) {
	size := len(s.min)
	if size == 0 {
		return 0, false
	}

	return s.min[size-1], true
}
