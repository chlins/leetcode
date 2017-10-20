package addTwoNumber

/**
 * https://leetcode.com/problems/add-two-numbers/description/
 *
 * You are given two non-empty linked lists representing two non-negative integers.
 * The digits are stored in reverse order and each of their nodes contain a single digit.
 * Add the two numbers and return it as a linked list.
 * You may assume the two numbers do not contain any leading zero, except the number 0 itself.
 *
 * Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
 * Output: 7 -> 0 -> 8
 */

// ListNode definition
type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * @param {*ListNode} l1
 * @param {*ListNode} l2
 * @return {*ListNode}
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	node := res
	node1 := l1
	node2 := l2
	over := false
	for node1 != nil || node2 != nil {
		tmp := 0
		if node1 != nil {
			tmp = tmp + node1.Val
			node1 = node1.Next
		}
		if node2 != nil {
			tmp = tmp + node2.Val
			node2 = node2.Next
		}
		if over {
			tmp++
		}
		if tmp >= 10 {
			over = true
			tmp = tmp % 10
		} else {
			over = false
		}
		node.Val = tmp
		if node1 != nil || node2 != nil {
			node.Next = &ListNode{}
			node = node.Next
		}
	}
	if over {
		node.Next = &ListNode{}
		node = node.Next
		node.Val = 1
	}
	return res
}
