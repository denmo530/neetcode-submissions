func groupAnagrams(strs []string) [][]string {
	groups := make(map[string][]string)

	for _, s := range strs {
		sRune := []rune(s)
		sort.Slice(sRune, func(i, j int) bool { 
			return sRune[i] < sRune[j] 
		})
		word := string(sRune)
		groups[word] = append(groups[word], s)
	}

	result := make([][]string, 0, len(groups))
	for _, group := range groups {
		result = append(result, group)
	}

	return result	
}
