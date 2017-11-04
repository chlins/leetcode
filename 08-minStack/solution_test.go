package minStack

import (
	"testing"
)

func TestMinStack(t *testing.T) {
	stack := Constructor()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	if stack.GetMin() != 1 {
		t.Errorf("The min is not right: %d\n", stack.GetMin())
	}
	stack.Pop()
	if stack.Top() != 2 {
		t.Errorf("The top is not right: %d\n", stack.Top())
	}
}
