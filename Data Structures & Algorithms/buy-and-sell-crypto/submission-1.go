func maxProfit(prices []int) int {
	left := 0
	maxProfit := 0

	for i := 1; i < len(prices); i++ {
		if prices[left] < prices[i] {
			currProfit := prices[i] - prices[left]
            maxProfit = int(math.Max(float64(maxProfit), float64(currProfit)))
		} else {
			left = i
		}
	} 

	return maxProfit
}
