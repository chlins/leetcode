package judgeRouteCircle

import (
	"testing"
)

/**
 * Test Cases:
 *   Case 1:
 *     Input: "UD"
 *     Except: true
 *   Case 2:
 *     Input: "LL"
 *     Expect: false
 */

func TestJudgeRouteCircle(t *testing.T) {
	input := "UD"
	if judgeCircle(input) {
		t.Log("### Case 1 passed!")
	} else {
		t.Error("### Case 1 failed!")
	}

	input = "LL"
	if !judgeCircle(input) {
		t.Log("### Case 2 passed!")
	} else {
		t.Error("### Case 2 failed!")
	}
}
