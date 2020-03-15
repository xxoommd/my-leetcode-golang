package leetcode

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isSymmetricTreesRecursion(root.Left, root.Right)
}

// isSymmetricTrees recursion method
func isSymmetricTreesRecursion(t1, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}

	if t1 == nil || t2 == nil {
		return false
	}

	if t1.Val != t2.Val {
		return false
	}

	return isSymmetricTreesRecursion(t1.Left, t2.Right) && isSymmetricTreesRecursion(t1.Right, t2.Left)
}

func isSymmetric2(t *TreeNode) bool {
	queue := SliceList{}
	queue.Push(t, t)

	for queue.Len() > 0 {
		t1, _ := queue.Shift().(*TreeNode)
		t2, _ := queue.Shift().(*TreeNode)

		if t1 == nil && t2 == nil {
			continue
		}
		if t1 == nil || t2 == nil {
			return false
		}
		if t1.Val != t2.Val {
			return false
		}

		queue.Push(t1.Left, t2.Right)
		queue.Push(t1.Right, t2.Left)
	}

	return true
}
