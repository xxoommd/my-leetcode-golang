package leetcode

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateTrees(t *testing.T) {
	type temp struct {
		n      int
		expect []*TreeNode
	}

	tests := []temp{
		temp{
			n:      0,
			expect: nil,
		},
		temp{
			n: 1,
			expect: []*TreeNode{
				NewTree(1),
			},
		},

		temp{
			n: 2,
			expect: []*TreeNode{
				NewTree(1, nil, NewTree(2)),
				NewTree(2, NewTree(1)),
			},
		},

		temp{
			n: 3,
			expect: []*TreeNode{
				NewTree(1, nil, NewTree(3, NewTree(2))),
				NewTree(3, NewTree(2, NewTree(1))),
				NewTree(3, NewTree(1, nil, NewTree(2))),
				NewTree(2, NewTree(1), NewTree(3)),
				NewTree(1, nil, NewTree(2, nil, NewTree(3))),
			},
		},
	}

	for _, test := range tests {
		got := generateTrees(test.n)
		sort.Slice(got, func(i, j int) bool { return got[i].Val < got[j].Val })
		sort.Slice(test.expect, func(i, j int) bool { return test.expect[i].Val < test.expect[j].Val })
		assert.Condition(t, func() bool {
			return isTreeNodeSliceEqual(got, test.expect)
		}, fmt.Sprintf("expect: %v\n   got: %v", test.expect, got))
	}
}
