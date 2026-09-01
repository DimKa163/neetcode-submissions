func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sort.Ints(nums)
	streak := 0
	result := 0
	for i := range len(nums) {
		if i > 0 {
			if nums[i] - nums[i - 1] > 1 {
				streak = 0
			}
			if nums[i] == nums[i-1] {
				continue
			}
		}
		streak++
		if streak > result {
			result = streak
		}
	}
	return result
}
