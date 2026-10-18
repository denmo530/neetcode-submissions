func topKFrequent(nums []int, k int) []int {
	counts := make(map[int]int)
	for _, n := range nums {
		counts[n]++
	}

	res := make([]int, 0, len(counts))

	for num := range counts {
    	res = append(res, num)
	}

	sort.Slice(res, func(i, j int) bool{
        return counts[res[j]] < counts[res[i]]
    })

	return res[:k]
}
