package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPreorderTraversal(t *testing.T) {
	input := NewTree(1, nil, NewTree(2, NewTree(3)))
	expect := []int{1, 2, 3}
	got := preorderTraversal(input)
	assert.Condition(t, func() bool {
		return isIntSliceEqual(expect, got)
	}, fmt.Sprintf("expect: %v\n   got: %v", expect, got))
}
