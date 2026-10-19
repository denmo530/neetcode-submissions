func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))
	p := 1
	
	for i:=0; i<len(nums); i++ {
		result[i] = p
		p *= nums[i]
	}
	p = 1
	for i:=len(nums)-1; i>=0; i-- {
		result[i] *= p
		p *= nums[i]
	}
	return result
}
