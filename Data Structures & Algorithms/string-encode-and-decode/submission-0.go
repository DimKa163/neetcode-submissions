

type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	var sb strings.Builder
	for _, str := range strs {
		lth := len(str)
		sb.WriteString(fmt.Sprintf("%d#%s", lth, str))
	}
	return sb.String()
}

func (s *Solution) Decode(encoded string) []string {
	strs := []string{}
	if encoded == "" {
		return strs
	}
	runes := []rune(encoded)
	i := 0
	for i < len(runes) {
		j := i
		for runes[j] != '#' {
			j++
		}
		lth, _ := strconv.Atoi(string(runes[i:j]))
		i = j + 1
		str := string(runes[i:i+lth])
		strs = append(strs, str)
		i += lth
	}
	return strs
}
