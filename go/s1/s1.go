package s1

func twoSum(nums []int, target int) []int {
	m := map[int]int{}

	for i, n := range nums {
		if v, ok := m[target-n]; ok && i != v {
			return []int{v, i}
		}

		m[n] = i
	}

	return nil
}
