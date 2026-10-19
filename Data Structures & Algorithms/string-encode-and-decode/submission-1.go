type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var b strings.Builder;

	for _, str := range strs {
		b.WriteString(strconv.Itoa(len(str)))
		b.WriteByte('#')
		b.WriteString(str)
	}

	return b.String()
}

func (s *Solution) Decode(encoded string) []string {
	var result []string
	i := 0

	for i < len(encoded) {
		n := 0
		for encoded[i] != '#' {
 			n = n*10 + int(encoded[i]-'0')
            i++;
		}
		i++;

		temp := encoded[i : i+n]
		result = append(result, temp)

		i += n
	}

	return result
}
