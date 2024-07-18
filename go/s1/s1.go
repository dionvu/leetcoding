package s1

func twoSum(nums []int, target int) []int {
	m := map[int]int{}

	for i, n := range nums {
		m[n] = i
	}

	for i, n := range nums {
		if v, ok := m[target-n]; ok && i != v {
			return []int{i, v}
		}
	}

	return nil
}
