package main

import (
	"strconv"
	"strings"
)

func permute(a []int) [][]int {
	n := len(a)
	if n == 0 {
		return nil
	}

	result := make([][]int, 0, factorial(n))
	elementsMap := make(map[string]struct{})
	addElement := func() {
		hash := sliceHash(a)
		if _, ok := elementsMap[hash]; ok {
			return
		}

		tmp := make([]int, n)
		copy(tmp, a)
		result = append(result, tmp)
		elementsMap[hash] = struct{}{}
	}
	addElement()

	p := make([]int, 0, n)
	for i := range a {
		p = append(p, i)
	}

	i := 1
	for i < n {
		p[i]--

		j := i % 2 * p[i]
		a[j], a[i] = a[i], a[j]

		addElement()

		i = 1
		for i < n && p[i] == 0 {
			p[i] = i
			i++
		}
	}

	return result
}

func sliceHash(a []int) string {
	const delim = ","
	builder := strings.Builder{}
	for i := range a {
		builder.WriteString(strconv.Itoa(a[i]))
		builder.WriteString(delim)
	}

	return builder.String()
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}

	res := 1
	for i := n; i > 0; i-- {
		res *= i
	}

	return res
}
