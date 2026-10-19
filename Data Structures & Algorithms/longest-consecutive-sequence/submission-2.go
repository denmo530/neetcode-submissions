import "slices"
func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	slices.Sort(nums)
	seqLen := 1
	longest := 1

	for i:=0; i < len(nums) - 1; i++ {
		curr := nums[i]
		next := nums[i+1]
		if next == curr + 1 {
			seqLen++
		} else if next == curr {
			continue
		} else {
			if seqLen > longest {
				longest = seqLen
			}
			seqLen = 1
		}
	}

	if seqLen > longest {
		longest = seqLen
	}

	return longest
}
