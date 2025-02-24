package stack

import (
	"errors"
	"fmt"
)

type Stack struct {
	items []int
}

// Push adds an item to the top of the stack
func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

// Pop removes and returns the top item from the stack
func (s *Stack) Pop() (int, error) {

	if s.isEmpty() {
		return 0, errors.New("Stack is empty")
	}

	top := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return top, nil
}

// Peek return the top element without removing it
func (s *Stack) Peek() (int, error) {

	if s.isEmpty() {
		return 0, errors.New("stack is empty")
	}
	return s.items[len(s.items)-1], nil
}

// func to check if Stack is empty
func (s *Stack) isEmpty() bool {
	return len(s.items) == 0
}

// func to return the number of elements in stack
func (s *Stack) Size() int {
	return len(s.items)
}

func main() {
	s := Stack{}
	s.Push(10)
	fmt.Printf("The size of Stack is: %d", s.Size())
}
