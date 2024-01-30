package main

import (
	"fmt"
	"strconv"
	"strings"
)

func findFirstOccurrence(arr []int, target int) int {
	l, r := 0, len(arr)-1
	ans := -1
	for l <= r {
		mid := l + (r-l)/2
		if arr[mid] == target {
			ans = mid
			r = mid - 1
		} else if arr[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return ans
}

func arrayAtoi(arr []string) []int {
	res := []int{}
	for _, x := range arr {
		v, _ := strconv.Atoi(x)
		res = append(res, v)
	}
	return res
}

func splitWords(s string) []string {
	if s == "" {
		return []string{}
	}
	return strings.Split(s, " ")
}

func main() {
	//scanner := bufio.NewScanner(os.Stdin)
	//scanner.Scan()
	//arr := arrayAtoi(splitWords(scanner.Text()))
	//scanner.Scan()
	//target, _ := strconv.Atoi(scanner.Text())

	arr := []int{1, 3, 3, 3, 3, 6, 10, 10, 10, 100}
	target := 3

	res := findFirstOccurrence(arr, target)
	fmt.Println(res)
}
