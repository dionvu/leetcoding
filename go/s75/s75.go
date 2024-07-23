package s75

func sortColors(nums []int) {
	for i := range nums {
		w := i - 1
		for w >= 0 && nums[w] > nums[w+1] {
			nums[w], nums[w+1] = nums[w+1], nums[w]
			w--
		}
	}
}
