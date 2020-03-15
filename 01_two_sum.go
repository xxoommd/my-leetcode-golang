package leetcode

func twoSum(nums []int, target int) []int {
	if len(nums) < 2 {
		return nil
	}

	res := make(map[int]int, len(nums))

	for i, v := range nums {
		k, ok := res[v]
		if ok {
			return []int{k, i}
		}

		res[target-v] = i
	}

	return nil
}
