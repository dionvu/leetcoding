package s739

func dailyTemperatures(temperatures []int) []int {
	stack := []int{}
	idxs := []int{}
	res := []int{}

	for i := 0; i < len(temperatures); i++ {
		res = append(res, 0)
	}

	for i, t := range temperatures {
		if i == 0 {
			stack = append(stack, t)
			idxs = append(idxs, i)
			continue
		}

		for len(stack) > 0 && stack[len(stack)-1] < t {
			last := len(stack) - 1

			res[idxs[last]] = i - idxs[last]

			stack = stack[:last]
			idxs = idxs[:last]
		}

		stack = append(stack, t)
		idxs = append(idxs, i)
	}

	return res
}
