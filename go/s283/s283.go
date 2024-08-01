package s283

func moveZeroes(nums []int) {
	// Will stay at the zero element
	i := 0

	for t := 0; t < len(nums); t++ {
		if nums[t] != 0 {
			nums[i], nums[t] = nums[t], nums[i]
		}

		if nums[i] != 0 {
			i++
		}
	}
}
