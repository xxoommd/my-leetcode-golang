package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLevelOrder(t *testing.T) {
	assert.Nil(t, levelOrder(nil), "Nil validation")

	input := NewTree(3, NewTree(9), NewTree(20, NewTree(15), NewTree(7)))
	expect := [][]int{
		[]int{3},
		[]int{9, 20},
		[]int{15, 7},
	}
	got := levelOrder(input)
	assert.Condition(t, func() bool {
		if len(expect) != len(got) {
			return false
		}
		for i := range expect {
			if !isIntSliceEqual(expect[i], got[i]) {
				return false
			}
		}
		return true
	}, fmt.Sprintf("expect: %v\n   got: %v", expect, got))
}
