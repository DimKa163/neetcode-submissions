func productExceptSelf(nums []int) []int {
	output := make([]int, len(nums))
	cur := 1
	pref := 1
	for i := 0; i < len(nums);i++ {
		output[i] = pref * cur
		cur = nums[i]
		pref = output[i]
	}
	postfix := 1
	for i := len(nums) - 1; i >= 0; i-- {
		
		output[i] = postfix * output[i] 
		postfix = nums[i] * postfix
	}
	return output
}
