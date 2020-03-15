package leetcode

func postorderTraversal(root *TreeNode) []int {
	var res []int

	stack := SliceList{}
	stack.Push(root)

	trav := SliceList{}

	var p *TreeNode

	for p != nil || stack.Len() > 0 {
		p, _ = stack.Pop().(*TreeNode)
		for p != nil {
			// cb(p)
			trav.Push(p)

			if p.Left != nil {
				stack.Push(p.Left)
			}
			p = p.Right
		}
	}

	for {
		p, _ = trav.Pop().(*TreeNode)
		if p == nil {
			break
		}
		res = append(res, p.Val)
	}
	return res
}
