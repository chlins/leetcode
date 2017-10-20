package addTwoNumber

import (
	"fmt"
	"reflect"
	"testing"
)

/**
 * Cases:
 *   1. Input:  2-->3-->4, 1-->3-->2
 *      Expect: 3-->6-->6
 *   2. Input:  2-->4-->3, 5-->6-->4
 *      Expect: 7-->0-->8
 *   3. Input:  7-->1-->6, 4-->9-->8
 *      Except: 1-->1-->5-->1
 */

func TestAddTwoNumber(t *testing.T) {
	// case one:
	inputL1 := &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}
	inputL2 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 2}}}
	expect := []int{3, 6, 6}
	output := make([]int, 3)
	node := addTwoNumbers(inputL1, inputL2)
	output[0] = node.Val
	output[1] = node.Next.Val
	output[2] = node.Next.Next.Val
	if reflect.DeepEqual(expect, output) {
		t.Log("### Case 1: passed!")
	} else {
		t.Error("### Case 1: failed!")
		fmt.Println(output)
	}

	// case two:
	inputL1 = &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}
	inputL2 = &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}
	expect = []int{7, 0, 8}
	node = addTwoNumbers(inputL1, inputL2)
	output[0] = node.Val
	output[1] = node.Next.Val
	output[2] = node.Next.Next.Val
	if reflect.DeepEqual(expect, output) {
		t.Log("### Case 2: passed!")
	} else {
		t.Error("### Case 2: failed!")
	}

	// case three:
	inputL1 = &ListNode{Val: 7, Next: &ListNode{Val: 1, Next: &ListNode{Val: 6}}}
	inputL2 = &ListNode{Val: 4, Next: &ListNode{Val: 9, Next: &ListNode{Val: 8}}}
	expect = []int{1, 1, 5, 1}
	node = addTwoNumbers(inputL1, inputL2)
	output1 := make([]int, 4)
	output1[0] = node.Val
	output1[1] = node.Next.Val
	output1[2] = node.Next.Next.Val
	output1[3] = node.Next.Next.Next.Val
	if reflect.DeepEqual(expect, output1) {
		t.Log("### Case 3: passed!")
	} else {
		t.Error("### Case 3: failed!")
	}
}
