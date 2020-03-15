package leetcode

func preorderTraversal(t *TreeNode) []int {
	var res []int

	stack := SliceList{}
	stack.Push(t)

	var p *TreeNode

	for p != nil || stack.Len() > 0 {
		v := stack.Pop()
		var ok bool
		p, ok = v.(*TreeNode)
		if ok {
			for p != nil {
				res = append(res, p.Val)

				if p.Right != nil {
					stack.Push(p.Right)
				}
				p = p.Left
			}
		}
	}
	return res
}
