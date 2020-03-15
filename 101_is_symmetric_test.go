package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSymmetric(t *testing.T) {
	tests := []struct {
		t      *TreeNode
		expect bool
	}{
		{nil, true},
		{NewTree(1), true},
		{
			NewTree(1, NewTree(2), NewTree(2)),
			true,
		},
		{
			NewTree(1, NewTree(2, nil, NewTree(3)), NewTree(2, nil, NewTree(3))),
			false,
		},
	}

	for _, test := range tests {
		got := isSymmetric(test.t)
		got2 := isSymmetric2(test.t)
		assert.Equal(t, test.expect, got)
		assert.Equal(t, test.expect, got2)
	}
}
