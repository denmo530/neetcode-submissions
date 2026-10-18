func hasDuplicate(nums []int) bool {
	values := make(map[int]int)

	for _, v := range nums {
		values[v]++
		if values[v] > 1 {
			return true
		}
	}

	return false
}
