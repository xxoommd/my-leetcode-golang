package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func isIntSliceEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func TestTwoSum(t *testing.T) {
	type input struct {
		target int
		nums   []int
		expect []int
	}

	nilTest := []input{
		input{0, []int{}, nil},
		input{0, []int{1}, nil},
		input{1, []int{1, 2}, nil},
		input{1, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, nil},
	}

	for _, input := range nilTest {
		assert.Nil(t, twoSum(input.nums, input.target))
	}

	nonNilTest := []input{
		input{0, []int{1, -1}, []int{0, 1}},
		input{9, []int{2, 3, 7, 5}, []int{0, 2}},
		input{11, []int{13, 24, 28, -1, -13}, []int{1, 4}},
	}

	for _, input := range nonNilTest {
		expect := input.expect
		got := twoSum(input.nums, input.target)

		assert.Condition(t, func() bool {
			return isIntSliceEqual(expect, got)
		}, fmt.Sprintf("expect: %v, got: %v", expect, got))
	}

	// expect2 := []int{1, 2}
	// got := twoSum([]int{0, 1, 2}, 3)
	// assert.Condition(t, func() bool {
	// 	return isIntSliceEqual(expect2, got)
	// }, fmt.Sprintf("expect: %v, got: %v", expect2, got))
	// assert.Equal(t, []int{1, 2}, twoSum([]int{0, 1, 2}, 3), "")
}
