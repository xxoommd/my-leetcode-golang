package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInorderTraversal(t *testing.T) {
	type input struct {
		biT    *TreeNode
		expect []int
	}

	tests := []input{
		input{
			biT: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val:  2,
					Left: &TreeNode{Val: 3},
				},
			},
			expect: []int{1, 3, 2},
		},
	}

	for _, test := range tests {
		got1 := inorderTraversal1(test.biT)
		assert.Condition(t, func() bool {
			return isIntSliceEqual(got1, test.expect)
		}, fmt.Sprintf("method 1: expect: %v, got: %v", test.expect, got1))

		got2 := inorderTraversal2(test.biT)
		assert.Condition(t, func() bool {
			return isIntSliceEqual(got2, test.expect)
		}, fmt.Sprintf("method 2: expect: %v, got: %v", test.expect, got2))
	}

}
