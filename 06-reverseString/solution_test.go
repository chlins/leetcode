package reverseString

import (
	"testing"
)

/**
 * Test Cases:
 *   Case 1:
 *     Input: "abcde"
 *     Except: "edcba"
 *
 *   Case 2:
 *     Input: "xyz zzy"
 *     Except: "yzz zyx"
 */

func TestReverseString(t *testing.T) {
	input := "abcde"
	if reverseString(input) != "edcba" {
		t.Error("### Case 1 failed!")
	} else {
		t.Log("### Case 1 passed!")
	}

	input = "xyz zzy"
	if reverseString(input) != "yzz zyx" {
		t.Error("### Case 2 failed!")
	} else {
		t.Log("### Case 2 passed!")
	}
}
