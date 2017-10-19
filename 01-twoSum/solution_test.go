package twoSum

import (
	"reflect"
	"testing"
)

/**
 * 1. input: [3,4,5] 7
 *    should output: [0,1]
 * 2. input: [1,4,5,9,2,7] 16
 *    should output: [3,5]
 * 3. input: [1,2,3] 10
 *    should output: [0,0]
 */

func TestTwoSum(t *testing.T) {
	inputOne := []int{3, 4, 5}
	outputOne := []int{0, 1}
	if reflect.DeepEqual(twoSum(inputOne, 7), outputOne) {
		t.Log("### Case 1: passed!")
	} else {
		t.Error("### Case 1: failed!")
	}

	inputTwo := []int{1, 4, 5, 9, 2, 7}
	outputTwo := []int{3, 5}
	if reflect.DeepEqual(twoSum(inputTwo, 16), outputTwo) {
		t.Log("### Case 2: passed!")
	} else {
		t.Error("### Case 2: failed!")
	}

	inputThree := []int{1, 2, 3}
	outputThree := []int{0, 0}
	if reflect.DeepEqual(twoSum(inputThree, 10), outputThree) {
		t.Log("### Case 3: passed!")
	} else {
		t.Error("### Case 3: failed!")
	}
}
