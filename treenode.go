package leetcode

import (
	"bytes"
	"fmt"
)

// TreeNode definition
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// traverseNode used for TreeNode traversal
type traverseNode func(node *TreeNode)

// NewTree is a handy way to create a new TreeNode. Args 'children' length should <= 2, more else will be ignored.
func NewTree(value int, children ...*TreeNode) *TreeNode {
	var l, r *TreeNode
	if len(children) > 0 {
		l = children[0]
	}
	if len(children) > 1 {
		r = children[1]
	}

	return &TreeNode{Val: value, Left: l, Right: r}
}

// PreorderTraverse traverses from root -> left child -> right child
func (t *TreeNode) PreorderTraverse(cb traverseNode) {
	// Recursion:
	/*
		if t == nil {
			return
		}

		cb(t) // traverse root by callback
		if t.Left != nil {
			t.Left.PreorderTraverse(cb)
		}
		if t.Right != nil {
			t.Right.PreorderTraverse(cb)
		}
	*/

	// Non-recursion: use Stack
	stack := SliceList{}
	stack.Push(t)

	var p *TreeNode

	for p != nil || stack.Len() > 0 {
		v := stack.Pop()
		var ok bool
		p, ok = v.(*TreeNode)
		if ok {
			for p != nil {
				cb(p)

				if p.Right != nil {
					stack.Push(p.Right)
				}
				p = p.Left
			}
		}
	}
}

// InorderTraverse traverses from left child -> root -> right child
func (t *TreeNode) InorderTraverse(cb traverseNode) {
	// Recursion:
	/*
		if t == nil {
			return nil
		}

		if t.Left != nil {
			t.Left.InorderTraverse(cb)
		}
		cb(t)
		if t.Right != nil {
			t.Right.InorderTraverse(cb)
		}
	*/

	// Non-recursion: use Stack
	stack := SliceList{}

	for t != nil || stack.Len() > 0 {
		for t != nil {
			stack.Push(t)
			t = t.Left
		}
		t, _ = stack.Pop().(*TreeNode)
		if t != nil {
			cb(t)
			t = t.Right
		}
	}
}

// PostorderTraverse traverses from left child -> right child -> root
func (t *TreeNode) PostorderTraverse(cb traverseNode) {
	// Recursion:
	/*
		if t == nil {
			return
		}

		if t.Left != nil {
			t.Left.PostorderTraverse(cb)
		}
		if t.Right != nil {
			t.Right.PostorderTraverse(cb)
		}
		cb(t)
	*/

	// use stack: similar as 'preorder' (always push right child to stack & reverse traversal order)
	stack := SliceList{}
	stack.Push(t)

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
		cb(p)
	}
}

// LevelOrderTraverse traverses from root, then children at the same level, then children's children at the same level...
func (t *TreeNode) LevelOrderTraverse(cb traverseNode) {
	queue := SliceList{}
	queue.Push(t)

	for {
		p, _ := queue.Shift().(*TreeNode)
		if p == nil {
			break
		}
		cb(p)
		if p.Left != nil {
			queue.Push(p.Left)
		}
		if p.Right != nil {
			queue.Push(p.Right)
		}
	}
}

func (t *TreeNode) hasChildren() bool {
	return t.Left != nil || t.Right != nil
}

func (t *TreeNode) noChild() bool {
	return t.Left == nil && t.Right == nil
}

func (t *TreeNode) String() string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d", t.Val))

	if t.hasChildren() {
		if t.Left != nil {
			buf.WriteString(fmt.Sprintf("->{%s", t.Left.String()))
		} else {
			buf.WriteString("->{nil")
		}
		if t.Right != nil {
			buf.WriteString(fmt.Sprintf(", %s}", t.Right.String()))
		} else {
			buf.WriteString(", nil}")
		}
	}
	return buf.String()
}

func compareTreeNodes(a, b *TreeNode) bool {
	aNil, bNil := a == nil, b == nil

	// both are nil
	if aNil && bNil {
		return true
	}

	// one of them is nil, the other is non-nil
	if (!aNil && bNil) || (aNil && !bNil) {
		return false
	}

	if a.Val != b.Val {
		return false
	}

	// comparing children
	if compareTreeNodes(a.Left, b.Left) && compareTreeNodes(a.Right, b.Right) {
		return true
	}

	return false
}

func isTreeNodeSliceEqual(a, b []*TreeNode) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if !compareTreeNodes(a[i], b[i]) {
			return false
		}
	}

	return true
}

func copyTree(tree *TreeNode) *TreeNode {
	if tree == nil {
		return nil
	}

	newTree := &TreeNode{Val: tree.Val}
	if tree.Left != nil {
		newTree.Left = copyTree(tree.Left)
	}
	if tree.Right != nil {
		newTree.Right = copyTree(tree.Right)
	}

	return newTree
}
