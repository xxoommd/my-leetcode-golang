package leetcode

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}

	pre := []*TreeNode{nil}

	for i := 1; i <= n; i++ {
		var curr []*TreeNode
		for _, preTree := range pre {
			// 1. add i as root
			newTree := &TreeNode{Val: i, Left: preTree}
			curr = append(curr, newTree)

			if preTree == nil {
				continue
			}

			// 2. if there is any right child exist: insert i as right child and make origin right child to left child
			//  else: just insert i as right child
			parent := preTree
			r := parent.Right
			for {
				parent.Right = &TreeNode{Val: i, Left: r}
				newTree := copyTree(preTree)
				curr = append(curr, newTree)
				parent.Right = r // reset preTree to origin value

				if r == nil {
					break
				}
				parent = r
				r = r.Right
			}
		}

		pre = curr
	}

	return pre
}
