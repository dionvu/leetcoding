package s75

// func sortColors(nums []int) {
// 	for i := range nums {
// 		w := i - 1
// 		for w >= 0 && nums[w] > nums[w+1] {
// 			nums[w], nums[w+1] = nums[w+1], nums[w]
// 			w--
// 		}
// 	}
// }

func sortColors(nums []int) {
	m := map[int]int{}

	for _, n := range nums {
		m[n]++
	}

	zeros := m[0]
	ones := m[1]
	twos := m[2]

	for i := 0; i < zeros; i++ {
		nums[i] = 0
	}

	for i := 0; i < ones; i++ {
		nums[i+zeros] = 1
	}

	for i := 0; i < twos; i++ {
		nums[i+zeros+ones] = 2
	}
}
