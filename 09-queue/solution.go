package queue

// https://leetcode.com/problems/implement-queue-using-stacks/description/
// Implement the following operations of a queue using stacks.
// push(x) -- Push element x to the back of queue.
// pop() -- Removes the element from in front of queue.
// peek() -- Get the front element.
// empty() -- Return whether the queue is empty.

type MyQueue struct {
	queue []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		queue: make([]int, 0),
	}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.queue = append(this.queue, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	defer func() {
		this.queue = this.queue[1:]
	}()
	return this.queue[0]
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	return this.queue[0]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.queue) == 0
}
