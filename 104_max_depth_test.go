package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxDepth(t *testing.T) {
	tests := []struct {
		t      *TreeNode
		expect int
	}{
		{nil, 0},
		{NewTree(1), 1},
		{NewTree(1, NewTree(2, NewTree(3), NewTree(4)), NewTree(2)), 3},
	}

	for _, test := range tests {
		assert.Equal(t, test.expect, maxDepth(test.t))
	}
}
