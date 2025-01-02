package hindex

import "sort"

func hIndex(citations []int) int {
	sort.Slice(citations, func(i, j int) bool {
		return citations[i] > citations[j]
	})

	i := 0
	for ; i < len(citations); i++ {
		if citations[i] < i+1 {
			return i
		}
	}

	return i
}
