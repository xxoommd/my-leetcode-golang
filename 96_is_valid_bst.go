package leetcode

import (
	"math"
)

func isValidBST(t *TreeNode) bool {
	return isValidBSTStack(t)
}

func isValidBSTRecursion(t *TreeNode) bool {
	if t == nil {
		return true
	}

	if t.Left != nil && t.Right != nil && t.Left.Val > t.Right.Val {
		return false
	}

	if t.Left != nil {
		if t.Left.Val >= t.Val || !isValidBSTRecursion(t.Left) {
			return false
		}
	}

	if t.Right != nil {
		if t.Right.Val <= t.Val || !isValidBSTRecursion(t.Right) {
			return false
		}
	}

	return true // t.Left == nil && t.Right == nil
}

func isValidBSTStack(t *TreeNode) bool {
	stack := SliceList{}
	last := math.MinInt64

	for t != nil || stack.Len() > 0 {
		for t != nil {
			stack.Push(t)
			t = t.Left
		}
		t, _ = stack.Pop().(*TreeNode)
		if t != nil {
			if t.Val <= last {
				return false
			}
			last = t.Val
			t = t.Right
		}
	}
	return true
}
