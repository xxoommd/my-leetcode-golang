package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSameTree(t *testing.T) {
	tests := []struct {
		t1     *TreeNode
		t2     *TreeNode
		expect bool
	}{
		{nil, nil, true},
		{
			NewTree(1, NewTree(2), NewTree(3)),
			NewTree(1, NewTree(2), NewTree(3)),
			true,
		},
		{
			NewTree(1, NewTree(2)),
			NewTree(1, nil, NewTree(2)),
			false,
		},
	}

	_ = tests
	for _, test := range tests {
		assert.Equal(t, test.expect, isSameTree(test.t1, test.t2))
	}
}
