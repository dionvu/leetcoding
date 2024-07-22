package s238

func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	for i := range nums {
		res[i] = 1
	}

	left := 1
	right := 1

	for i, n := range nums {
		res[i] *= left
		left *= n
	}

	for i := len(nums) - 1; i >= 0; i-- {
		res[i] *= right
		right *= nums[i]
	}

	return res
}
