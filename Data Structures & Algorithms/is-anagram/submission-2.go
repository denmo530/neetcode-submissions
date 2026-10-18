func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	
	mapS, mapT := make(map[rune]int), make(map[rune]int)
	for i, char := range s {
		mapS[char]++
		mapT[rune(t[i])]++
	}

	for i, v := range mapS {
		if mapT[i] != v {
			return false
		}
	}

	return true
}