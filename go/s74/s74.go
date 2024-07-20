package s74

import "sort"

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || target < matrix[0][0] {
		return false
	}

	for _, vec := range matrix {
		if target < vec[0] || vec[len(vec)-1] < target {
			continue
		}

		idx := sort.Search(len(vec), func(i int) bool { return target <= vec[i] })

		return idx < len(vec) && vec[idx] == target
	}

	return false
}
