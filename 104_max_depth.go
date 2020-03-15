package leetcode

func maxDepth(t *TreeNode) int {
	// return maxDepthRecursion(t)
	return maxDepth2(t)
}

func maxDepthRecursion(t *TreeNode) int {
	if t == nil {
		return 0
	}

	return 1 + maxInt(maxDepth(t.Left), maxDepth(t.Right))
}

func maxDepth2(t *TreeNode) int {
	if t == nil {
		return 0
	}

	type levelTree struct {
		tree  *TreeNode
		level int
	}

	stack := SliceList{}
	stack.Push(&levelTree{t, 1})

	var res int

	for stack.Len() > 0 {
		curr, _ := stack.Pop().(*levelTree)
		res = maxInt(res, curr.level)

		if curr.tree.Left != nil {
			stack.Push(&levelTree{curr.tree.Left, curr.level + 1})
		}
		if curr.tree.Right != nil {
			stack.Push(&levelTree{curr.tree.Right, curr.level + 1})
		}
	}

	return res
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
