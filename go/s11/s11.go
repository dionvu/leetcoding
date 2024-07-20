package s11

func maxArea(height []int) int {
	left := 0
	right := len(height) - 1

	max_area := 0

	for left < right {
		curr := min(height[left], height[right]) * (right - left)

		max_area = max(curr, max_area)

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return max_area
}
