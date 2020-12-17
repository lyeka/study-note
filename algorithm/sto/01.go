package sto

func findRepeatNumber(nums []int) int {
	for i, v := range nums {
		for i != v {
			if v == nums[v] {
				return v
			}
			nums[v], nums[i] = nums[i], nums[v]
		}
	}

	return -1
}

func findRepeatNumber2(nums []int) int {
	tmp := make([]int, len(nums))
	for _, v := range nums {
		tmp[v]++
		if tmp[v] > 1 {
			return v
		}
	}
	return -1
}