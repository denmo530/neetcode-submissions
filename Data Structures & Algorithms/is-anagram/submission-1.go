func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	
	mapS := make(map[byte]int, len(s))
	mapT := make(map[byte]int, len(t))

	for _, v := range []byte(s) {
		mapS[v]++
	}
	for _, v := range []byte(t) {
		mapT[v]++
	}

	for char, count := range mapS {
    otherCount, ok := mapT[char]

    if !ok || count != otherCount {
		return false
    	}
	}

	return true
}