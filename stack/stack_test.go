package stack

import "testing"

func TestStack_Push(t *testing.T) {
	stack := Stack{}

	stack.Push(10)
	stack.Push(20)

	val, err := stack.Pop()
	if err != nil || val != 20 {
		t.Errorf("Expected 20, got %d", val)
	}

	val, err = stack.Pop()
	if err != nil || val != 10 {
		t.Errorf("Expected 10, got %d", val)
	}

	if !stack.isEmpty() {
		t.Errorf("Extpected stack to be empty")
	}

}
