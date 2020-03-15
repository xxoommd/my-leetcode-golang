package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHasChildren(t *testing.T) {
	a := &TreeNode{
		Val:  1,
		Left: &TreeNode{Val: 2},
	}
	assert.Equal(t, a.hasChildren(), true, "")

	b := &TreeNode{Val: 1}
	assert.Equal(t, b.hasChildren(), false, "")
}

func TestCompareTreeNode(t *testing.T) {
	a := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2, Right: &TreeNode{Val: 3}},
		Right: &TreeNode{Val: 2},
	}
	b := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2, Right: &TreeNode{Val: 3}},
		Right: &TreeNode{Val: 2},
	}
	c := &TreeNode{
		Val:   1,
		Right: &TreeNode{Val: 2},
	}

	var d *TreeNode

	assert.Equal(t, true, compareTreeNodes(a, b), "")
	assert.Equal(t, false, compareTreeNodes(a, c), "")
	assert.Equal(t, false, compareTreeNodes(a, d), "")
}

func TestTraversals(t *testing.T) {
	type input struct {
		tree        *TreeNode
		expectPre   []int
		expectIn    []int
		expectPost  []int
		expectLevel []int
	}

	tests := []input{
		input{
			tree:        nil,
			expectPre:   nil,
			expectIn:    nil,
			expectPost:  nil,
			expectLevel: nil,
		},
		input{
			tree:        NewTree(1),
			expectPre:   []int{1},
			expectIn:    []int{1},
			expectPost:  []int{1},
			expectLevel: []int{1},
		},
		input{
			tree:        NewTree(1, NewTree(2, NewTree(4), NewTree(5, NewTree(6), NewTree(7))), NewTree(3, nil, NewTree(8))),
			expectPre:   []int{1, 2, 4, 5, 6, 7, 3, 8},
			expectIn:    []int{4, 2, 6, 5, 7, 1, 3, 8},
			expectPost:  []int{4, 6, 7, 5, 2, 8, 3, 1},
			expectLevel: []int{1, 2, 3, 4, 5, 8, 6, 7},
		},
	}

	for _, test := range tests {
		var gotPre, gotIn, gotPost, gotLevel []int
		test.tree.PreorderTraverse(func(node *TreeNode) {
			gotPre = append(gotPre, node.Val)
		})
		test.tree.InorderTraverse(func(node *TreeNode) {
			gotIn = append(gotIn, node.Val)
		})
		test.tree.PostorderTraverse(func(node *TreeNode) {
			gotPost = append(gotPost, node.Val)
		})
		test.tree.LevelOrderTraverse(func(node *TreeNode) {
			gotLevel = append(gotLevel, node.Val)
		})
		assert.Condition(t, func() bool {
			return isIntSliceEqual(gotPre, test.expectPre)
		}, fmt.Sprintf("Preorder - expect: %v\ngot: %v", test.expectPre, gotPre))
		assert.Condition(t, func() bool {
			return isIntSliceEqual(gotIn, test.expectIn)
		}, fmt.Sprintf("Inorder - expect: %v\ngot: %v", test.expectIn, gotIn))
		assert.Condition(t, func() bool {
			return isIntSliceEqual(gotPost, test.expectPost)
		}, fmt.Sprintf("Postorder - expect: %v\ngot: %v", test.expectPost, gotPost))
		assert.Condition(t, func() bool {
			return isIntSliceEqual(gotLevel, test.expectLevel)
		}, fmt.Sprintf("Level-order - expect: %v\ngot: %v", test.expectLevel, gotLevel))
	}
}
