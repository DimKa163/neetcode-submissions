func twoSum(nums []int, target int) []int {
    m := make(map[int]int)
	result := make([]int, 2)
	for i, v := range nums {
		m[v] = i
	}
	for i, v := range nums {
		a, ok := m[target - v]
		if !ok || a == i {
			continue
		}
		result[0] = i
		result[1] = a
		return result
	}
	return result
}
