func lengthOfLongestSubstring(s string) int {
	l := 0
	maxLen := 0
	seen := make(map[byte]int)

	for r := 0; r < len(s); r++ {
		char := s[r]

		if prevIndex, ok := seen[char]; ok && prevIndex >= l {
			l = prevIndex + 1
		}

		seen[char] = r
		windowLen := r - l + 1

		if windowLen > maxLen {
				maxLen = windowLen
		}
	}

	return maxLen
}
