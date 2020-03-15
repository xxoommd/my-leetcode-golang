package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func is2dSliceEquil(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if !isIntSliceEqual(a[i], b[i]) {
			return false
		}
	}

	return true
}

func TestThreeSum(t *testing.T) {
	type input struct {
		nums   []int
		expect [][]int
	}

	tests := []input{
		input{
			nums: []int{-1, 0, 1, 2, -1, -4},
			expect: [][]int{
				[]int{-1, -1, 2},
				[]int{-1, 0, 1},
			},
		},
		input{
			nums: []int{1, 1, 1, 0, 0, 0, 0},
			expect: [][]int{
				[]int{0, 0, 0},
			},
		},
	}

	for _, test := range tests {
		got := threeSum(test.nums)
		assert.Condition(t, func() bool {
			return is2dSliceEquil(test.expect, got)
		}, fmt.Sprintf("expect %v, got %v", test.expect, got))
	}
}
