func maxArea(heights []int) int {
	highest := 0

	left := 0
	right := len(heights) - 1

	for left < right {
		minH := int(math.Min(float64(heights[left]), float64(heights[right])))
		width := right - left
		max := minH * width

		if max > highest {
			highest = max
		}

		if heights[left] < heights[right] {
			left++
		} else {
			right--
		}
	}

	return highest
}
