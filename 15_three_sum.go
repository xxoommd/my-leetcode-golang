package leetcode

import (
	"sort"
)

// 解法一：暴力破解，三重循环；可按twoSum()的方式进行优化
// 解法二：
// 	1. 数组进行升序排序
// 	2. 遍历排序后的数组。因升序排列，如果遍历到的元素>0，则之后无答案，结束。
// 	3. 假设每次遍历的数组索引为index，设置两个数组指针i和j，i=index+1, j=数组最后一位，i++，j--（向i和j的中间靠拢）。
// 	三数之和为0，即答案。
// 	4. 去重的关键在于i和j去重。如果 i+1 == i，则跳过i++；同理 j-1 == j，也跳过j--。
func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}

	// Do not want to change the origin slice
	arr := make([]int, len(nums))
	copy(arr, nums)
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	var result [][]int

	for index, num := range arr {
		if num > 0 {
			break
		}

		if index > 0 && num == arr[index-1] { // skip the previous same value
			continue
		}

		i, j := index+1, len(arr)-1
		for i < j {
			v := num + arr[i] + arr[j]
			if v == 0 { // bingo
				if result == nil {
					result = [][]int{}
				}
				result = append(result, []int{num, arr[i], arr[j]})
				for i < j && arr[i] == arr[i+1] {
					i++
				}
				for i < j && arr[j] == arr[j-1] {
					j--
				}
				i++
				j--
			} else if v > 0 {
				j--
			} else /* if v < 0 */ {
				i++
			}
		}
	}

	return result
}
