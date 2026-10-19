type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var encoded string;

	for _, str := range strs {
		encoded += strconv.Itoa(len(str))
		encoded += "#"
		encoded += str
	}

	return encoded
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
