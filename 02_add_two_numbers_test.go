package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddTwoNumbers(t *testing.T) {
	assert.Nil(t, addTwoNumbers(nil, nil))

	type input struct {
		l1     *ListNode
		l2     *ListNode
		expect *ListNode
	}

	nonNilTests := []input{
		input{
			CreateListNode(2, 4, 3),
			CreateListNode(5, 6, 4),
			CreateListNode(7, 0, 8),
		},
		input{
			CreateListNode(5),
			CreateListNode(5),
			CreateListNode(0, 1),
		},
		input{
			CreateListNode(1, 8),
			CreateListNode(0),
			CreateListNode(1, 8),
		},
	}

	for _, test := range nonNilTests {
		got := addTwoNumbers(test.l1, test.l2)
		assert.Condition(t, func() bool {
			return isListNodeEqual(test.expect, got)
		}, fmt.Sprintf("input: %v, %v\nexpect: %v, got: %v", test.l1, test.l2, test.expect, got))
	}
}
