package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostorderTraversal(t *testing.T) {
	input := NewTree(1, nil, NewTree(2, NewTree(3)))
	expect := []int{3, 2, 1}
	got := postorderTraversal(input)
	assert.Condition(t, func() bool {
		return isIntSliceEqual(expect, got)
	}, fmt.Sprintf("expect: %v\n got: %v", expect, got))
}
