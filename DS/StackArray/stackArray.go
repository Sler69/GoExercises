package stack

type stack struct {
	length      int
	stackMemory []interface{}
	index       int
}

func (s *stack) Push(value interface{}) {
	s.stackMemory[s.index] = value
	s.length++
	s.index++
	return
}

func (s *stack) Pop() interface{} {
	if s.length == 0 {
		return nil
	}
	var topValue interface{} = s.stackMemory[s.length-1]
	s.stackMemory[s.length-1] = nil
	s.length--
	s.index--
	return topValue

}

func (s *stack) IsEmpty() bool {
	return s.length == 0
}

func (s *stack) Peek() interface{} {
	return s.stackMemory[s.length+1]
}

// New creates a stack array
func New(size int) stack {
	if size == 0 {
		size = 1000
	}
	s := stack{0, make([]interface{}, size), 0}
	return s
}
