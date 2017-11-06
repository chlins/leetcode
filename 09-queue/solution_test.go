package queue

import (
	"testing"
)

func TestQueue(t *testing.T) {
	queue := Constructor()
	queue.Push(1)
	queue.Push(2)
	queue.Push(3)
	if queue.Peek() != 1 {
		t.Error("### Test failed!")
	}
	queue.Pop()
	if queue.Peek() != 2 {
		t.Error("### Test failed!")
	}
	if queue.Empty() {
		t.Error("### Test failed!")
	}
}
