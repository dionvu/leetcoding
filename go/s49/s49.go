package s49

import (
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	groups := map[string][]string{}
	result := [][]string{}

	for _, s := range strs {
		sp := strings.Split(s, "")
		sort.Strings(sp)

		so := strings.Join(sp, "")

		groups[so] = append(groups[so], s)
	}

	for _, g := range groups {
		result = append(result, g)
	}

	return result
}
