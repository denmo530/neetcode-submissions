func stripNonAlphanumeric(s string) string {
    return strings.Map(func(r rune) rune {
        if unicode.IsLetter(r) || unicode.IsDigit(r) {
            return r
        }
        return -1
    }, s)
}

func isPalindrome(s string) bool {
	filteredS := strings.ToLower(stripNonAlphanumeric(s))
	left := 0; 
	right := len(filteredS)-1

	for left < right {
		p1 := filteredS[left]
		p2 := filteredS[right]

		if p1 != p2 {
			return false
		}

		left++
		right--
	}
	return true
}
