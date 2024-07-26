package s347

import (
	"slices"
)

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for _, n := range nums {
		m[n]++
	}

	res := [][]int{}
	for n, v := range m {
		res = append(res, []int{n, v})
	}

	slices.SortFunc(res, func(a []int, b []int) int {
		return b[1] - a[1]
	})

	topK := make([]int, k)
	for i := 0; i < k; i++ {
		topK[i] = res[i][0]
	}

	return topK
}
