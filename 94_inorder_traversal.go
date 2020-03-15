package leetcode

// 递归实现 ez
func inorderTraversal1(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var res []int

	if root.Left != nil {
		res = append(res, inorderTraversal1(root.Left)...)
	}
	res = append(res, root.Val)
	if root.Right != nil {
		res = append(res, inorderTraversal1(root.Right)...)
	}

	return res
}

// 基于栈的遍历
func inorderTraversal2(root *TreeNode) []int {
	var res []int

	stack := SliceList{}

	t := root
	for t != nil || stack.Len() > 0 {
		for t != nil {
			stack.Push(t)
			t = t.Left
		}
		v := stack.Pop()
		var ok bool
		t, ok = v.(*TreeNode)
		if ok {
			res = append(res, t.Val)
			t = t.Right
		}
	}

	return res
}
