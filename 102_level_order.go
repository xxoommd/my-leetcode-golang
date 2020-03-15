package leetcode

func levelOrder(root *TreeNode) [][]int {
	// return levelOrderRecursion(root, 0, nil)

	return levelOrderQueue(root)
}

// levelOrderRecursion 递归实现
// - 递归很好理解，但我觉得是个伪层次遍历，因为其遍历的时候是深度遍历，再将结果用level作为层级计数进行整合。
func levelOrderRecursion(t *TreeNode, level int, res [][]int) [][]int {
	if t == nil {
		return nil
	}

	if len(res) == level {
		res = append(res, []int{})
	}

	// fullfil current level
	res[len(res)-1] = append(res[len(res)-1], t.Val)

	if t.Left != nil {
		res = levelOrderRecursion(t.Left, level+1, res)
	}
	if t.Right != nil {
		res = levelOrderRecursion(t.Right, level+1, res)
	}
	return res
}

// levelOrderQueue 迭代实现
// - 该方法用到了queue。第一个for循环为循环每一层，第二个for循环为遍历该层所有的节点。
func levelOrderQueue(t *TreeNode) [][]int {
	if t == nil {
		return nil
	}

	var res [][]int

	queue := SliceList{}
	queue.Push(t)

	for queue.Len() > 0 {
		l := queue.Len()
		res = append(res, []int{})
		for i := 0; i < l; i++ {
			curr, _ := queue.Shift().(*TreeNode) // curr should not be nil
			res[len(res)-1] = append(res[len(res)-1], curr.Val)

			if curr.Left != nil {
				queue.Push(curr.Left)
			}
			if curr.Right != nil {
				queue.Push(curr.Right)
			}
		}
	}

	return res
}
