func twoSum(nums []int, target int) []int {
	numsMap := make(map[int]int, len(nums))

	for i, num := range nums {
    	numsMap[num] = i
	}

	for i := 0; i < len(nums); i++ {
		diff := target - nums[i]
		j, ok := numsMap[diff]; if ok && i != j {
			return []int{i, j}	
		}
	}
	return []int{}
}
