func characterReplacement(s string, k int) int {
	l := 0
	maxFreq := 0
	maxLen := 0
	freq := make(map[byte]int)

	for r := 0; r < len(s); r++ {
		char := s[r]

	freq[char]++
	windowLen := r - l + 1

	if freq[char] > maxFreq {
		maxFreq = freq[char]
	}

	if windowLen - maxFreq > k {
    	freq[s[l]]--
    	l++
	}

	windowLen = r - l + 1
	if windowLen > maxLen {
		maxLen = windowLen
	}
}
	return maxLen
}
