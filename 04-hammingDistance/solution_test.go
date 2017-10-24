package hammingDistance

import (
	"testing"
)

/**
 * Test Cases:
 *   Case 1:
 *     10 (1 0 1 0)
 *     8  (1 0 0 0)
 *     except: 1
 *
 *   Case 2:
 *     15 (1 1 1 1)
 *     0  (0 0 0 0)
 *     except: 4
 */

func TestHammingDistance(t *testing.T) {
	x1 := 10
	x2 := 8
	if hammingDistance(x1, x2) != 1 {
		t.Error("### Case 1 failed!")
	} else {
		t.Log("### Case 1 passed!")
	}

	x1 = 15
	x2 = 0
	if hammingDistance(x1, x2) != 4 {
		t.Error("### Case 2 failed!")
	} else {
		t.Log("### Case 2 passed!")
	}
}
