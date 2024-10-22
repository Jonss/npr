package npr

type stack struct {
	st []int
}

func NewStack() stack {
	return stack{[]int{}}
}

func (s *stack) Pop() int {
	latest := len(s.st) - 1
	value := s.st[latest]
	s.st = s.st[:latest]
	return value
}

func (s *stack) Push(n int) {
	s.st = append(s.st, n)
}

func (s *stack) Len() int {
	return len(s.st)
}
