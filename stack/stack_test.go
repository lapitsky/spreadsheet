package stack

import (
	"testing"

	"github.com/lapitsky/spreadsheet/cells"
)

func TestNewStack(t *testing.T) {
	stack := NewStack()
	if stack.Len() != 0 {
		t.Errorf("Len for a new stack was incorrect, got: %d, want: %d.", stack.Len(), 0)
	}
}

func TestStack_PushAndPop(t *testing.T) {
	one := cells.NewFloatValue(1)
	two := cells.NewFloatValue(1)
	three := cells.NewFloatValue(1)

	stack := NewStack()

	stack.Push(one)
	stack.Push(two)
	stack.Push(three)

	if stack.Len() != 3 {
		t.Errorf("Len for a stack was incorrect, got: %d, want: %d.", stack.Len(), 3)
	}

	if val, ok := stack.Pop(); !ok || val != three {
		t.Errorf("Expected 3 to be at the top of the stack")
	}
	if val, ok := stack.Pop(); !ok || val != two {
		t.Errorf("Expected 2 to be at the top of the stack")
	}
	if val, ok := stack.Pop(); !ok || val != one {
		t.Errorf("Expected 1 to be at the top of the stack")
	}
	if _, ok := stack.Pop(); ok {
		t.Errorf("Expected Pop() from an empty stack to be unsuccessful")
	}
}
