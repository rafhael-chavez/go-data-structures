package stack

import (
	"testing"
)

func TestStack(t *testing.T) {
	s := NewStack()

	// Push
	s.Push(10)
	s.Push(20)
	s.Push(30)

	// Pop
	val, err := s.Pop()
	if err != nil || val != 30 {
		t.Errorf("Expected 30, got %d", val)
	}

	// Vaciar stack
	s.Pop()
	s.Pop()

	// Intentar Pop en stack vacío
	_, err = s.Pop()
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}
