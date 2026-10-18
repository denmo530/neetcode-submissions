func groupAnagrams(strs []string) [][]string {
	groups := make(map[[26]int][]string)

	for _, s := range strs {
		var freq [26]int
		for _, char := range s {
			freq[char-'a']++
		}
		groups[freq] = append(groups[freq], s)
	}

	result := make([][]string, 0, len(groups))
	for _, group := range groups {
		result = append(result, group)
	}

	return result	
}
