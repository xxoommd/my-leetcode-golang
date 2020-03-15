package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValidBST(t *testing.T) {
	tests := []struct {
		t      *TreeNode
		expect bool
	}{
		{nil, true},
		{NewTree(1), true},
		{NewTree(1, NewTree(1)), false},
		{
			NewTree(2, NewTree(1), NewTree(3)),
			true,
		},
		{
			NewTree(5, NewTree(1), NewTree(4, NewTree(3), NewTree(6))),
			false,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expect, isValidBSTRecursion(test.t), fmt.Sprintf("Recursion: %v", test.t))
		assert.Equal(t, test.expect, isValidBSTStack(test.t), fmt.Sprintf("Stack: %v", test.t))
	}
}
