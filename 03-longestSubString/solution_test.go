package longestSubString

import (
	"testing"
)

/**
 * Test Cases:
 *   case 1:
 *     Input: "abcdea"
 *     Except: 5
 *   case 2:
 *     Input: "abcabcabcdef"
 *     Expect: 6
 *   case 3:
 *     Input: ""
 *     Expect: 0
 */

func TestLongestSubString(t *testing.T) {
	input1 := "abcdea"
	output1 := lengthOfLongestSubstring(input1)
	if output1 == 5 {
		t.Log("### Case 1 passed!")
	} else {
		t.Error("### Case 1 failed!")
	}

	input2 := "abcabcabcdef"
	output2 := lengthOfLongestSubstring(input2)
	if output2 == 6 {
		t.Log("### Case 2 passed!")
	} else {
		t.Error("### Case 2 failed!")
	}

	input3 := ""
	output3 := lengthOfLongestSubstring(input3)
	if output3 == 0 {
		t.Log("### Case 3 passed!")
	} else {
		t.Error("### Case 3 failed!")
	}
}
