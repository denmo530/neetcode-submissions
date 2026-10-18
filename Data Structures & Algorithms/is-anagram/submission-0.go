func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sBytes := []byte(s)
	tBytes := []byte(t)
	sort.Slice(sBytes, func(i, j int) bool { return sBytes[i] < sBytes[j] })
	sort.Slice(tBytes, func(i, j int) bool { return tBytes[i] < tBytes[j] })

	if string(sBytes) == string(tBytes) {
		return true
	}

	return false
}