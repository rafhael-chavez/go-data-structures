package stack

import "fmt"

type Stack struct {
	elements []int
}

func NewStack() *Stack {
	return &Stack{elements: []int{}}
}

func (s *Stack) Push(value int) {
	s.elements = append(s.elements, value)
}

func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, fmt.Errorf("stack vacío")
	}
	lastIndex := len(s.elements) - 1
	value := s.elements[lastIndex]
	s.elements = s.elements[:lastIndex]
	return value, nil
}

func (s *Stack) Peek() (int, error) {
	if s.IsEmpty() {
		return 0, fmt.Errorf("stack vacío")
	}
	return s.elements[len(s.elements)-1], nil
}

func (s *Stack) IsEmpty() bool {
	return len(s.elements) == 0
}

func (s *Stack) ShowStack() {
	fmt.Println("Stack actual:", s.elements)
}
