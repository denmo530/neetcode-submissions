func twoSum(nums []int, target int) []int {
    for i := 0; i < len(nums); i++ {
		for j := i+1; j < len(nums); j++ {
			candidate := nums[i] + nums[j]
			if candidate == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}
