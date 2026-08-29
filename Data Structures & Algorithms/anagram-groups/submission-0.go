import "slices"
func groupAnagrams(strs []string) [][]string {
	result := make([][]string, 0, len(strs) / 2 + 1)
	slices.SortFunc(strs, func(a, b string) int {
		ar := []rune(a)
		br := []rune(b)
		slices.Sort(ar)
		slices.Sort(br)
		return slices.Compare(ar, br)
	})
	sub := make([]string, 0)
	i := 0
	sub = append(sub, strs[i])
	i++
	for i < len(strs) {
		cur := []rune(strs[i])
		prv := []rune(strs[i - 1])
		slices.Sort(cur)
		slices.Sort(prv)
		if slices.Compare(cur, prv) == 0 {
			sub = append(sub, strs[i])
		} else {
			result = append(result, sub)
			sub = make([]string, 0)
			sub = append(sub, strs[i])
		}
		i++
	}
	result = append(result, sub)
	return result
}
