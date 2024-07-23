package s287

func findDuplicate(nums []int) int {
	slow, fast := nums[0], nums[0]

	// Find the intersection
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}

	// Distance for the start to the result, and distance
	// from the intersection to result is always the same
	slow2 := nums[0]
	for slow != fast {
		slow = nums[slow]
		slow2 = nums[slow2]
	}

	return slow
}

// https://www.youtube.com/watch?v=wjYnzkAhcNk
