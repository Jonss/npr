package npr

type stack struct {
	st []int
}

func NewStack() stack {
	return stack{[]int{}}
}

func (s *stack) Pop() int {
	value := s.st[len(s.st)-1]
	s.st = s.st[:len(s.st)-1]
	return value
}

func (s *stack) Push(n int) {
	s.st = append(s.st, n)
}
