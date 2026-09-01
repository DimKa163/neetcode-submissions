func longestConsecutive(nums []int) int {
	m := make(map[int]int)
	var result int
	for _, num := range nums {
		if m[num] != 0 {
			continue
		}
		l := m[num - 1]
		r := m[num + 1]
		sum := l + r + 1
		m[num] = sum
		m[num - l] = sum
		m[num + r] = sum
		if result < sum {
			result = sum
		}
	}
	return result
}
