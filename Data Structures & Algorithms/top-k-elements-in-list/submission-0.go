import "slices"
func topKFrequent(nums []int, k int) []int {
	counter := make(map[int]int)
	for _, num := range nums {
		counter[num] = counter[num] + 1
	}
	pairs := make([][2]int, 0)
	for k, v := range counter {
		pairs = append(pairs, [2]int{k, v})
	}
	slices.SortFunc(pairs, func(l, r [2]int) int {
		return r[1] - l[1]
	})
	result := make([]int, k)
	for i, pair := range pairs {
		if i >= k {
			break
		}
		result[i] = pair[0]
	}
	return result
}