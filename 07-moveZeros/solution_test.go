package moveZeros

import (
	"reflect"
	"testing"
)

func TestMoveZeros(t *testing.T) {
	input := []int{1, 0, 2, 3, 0, 7, 9}
	moveZeroes(input)
	if !reflect.DeepEqual(input, []int{1, 2, 3, 7, 9, 0, 0}) {
		t.Errorf("Output is not right: %v", input)
	}
}
