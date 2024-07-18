package s128

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	m := make(map[int]bool)
	for _, n := range nums {
		m[n] = true
	}

	longest_streak := 0

	for _, n := range nums {
		if m[n-1] {
			continue
		}

		cur := n
		streak := 1

		for m[cur+1] {
			cur++
			streak++
		}

		longest_streak = max(longest_streak, streak)
	}

	return longest_streak
}
