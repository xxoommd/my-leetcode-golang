package leetcode

import (
	"bytes"
	"fmt"
)

// ListNode is a singly-linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	p := l
	for p != nil {
		if p.Next == nil {
			buf.WriteString(fmt.Sprint(p.Val))
		} else {
			buf.WriteString(fmt.Sprintf("%d -> ", p.Val))
		}
		p = p.Next
	}
	buf.WriteByte(']')
	return buf.String()
}

// CreateListNode creates *ListNode by an int array
func CreateListNode(arr ...int) *ListNode {
	if len(arr) == 0 {
		return nil
	}

	list := &ListNode{}
	var l *ListNode

	for _, v := range arr {
		if l == nil {
			l = list
			l.Val = v
		} else {
			l.Next = &ListNode{Val: v}
			l = l.Next
		}
	}

	return list
}

func isListNodeEqual(l1, l2 *ListNode) bool {
	for {
		if l1 == nil && l2 == nil {
			return true
		}
		if l1 == nil && l2 != nil {
			return false
		}
		if l1 != nil && l2 == nil {
			return false
		}
		if l1.Val != l2.Val {
			return false
		}
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
}

func addTwoNumbers(l1, l2 *ListNode) *ListNode {
	var result, p *ListNode
	var carry int

	for l1 != nil || l2 != nil {
		var x, y int
		if l1 != nil {
			x = l1.Val
		}
		if l2 != nil {
			y = l2.Val
		}

		z := x + y + carry
		if z >= 10 {
			carry = 1
			z -= 10
		} else {
			carry = 0
		}

		if result == nil {
			result = &ListNode{Val: z}
			p = result
		} else {
			p.Next = &ListNode{Val: z}
			p = p.Next
		}

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}

	if carry > 0 {
		p.Next = &ListNode{Val: carry}
	}

	return result
}
