package main

import "fmt"

var factorials = make(map[int]int, 10)

func countSpecialNumbers(n int) int {
	nDigits := make([]int, 0, 10)
	for i := n + 1; i > 0; i /= 10 {
		nDigits = append(nDigits, i%10)
	}
	nLen := len(nDigits)

	res := 0
	for i := 0; i < nLen-1; i++ {
		res += 9 * perms(9, i)
	}

	foundDigit := make(map[int]struct{}, 10)
	for i := nLen - 1; i >= 0; i-- {
		digit := nDigits[i]
		start := 0
		if i == nLen-1 {
			start = 1
		}
		for j := start; j < digit; j++ {
			if _, found := foundDigit[j]; !found {
				res += perms(9+i-nLen+1, i)
			}
		}
		if _, found := foundDigit[digit]; found {
			break
		}
		foundDigit[digit] = struct{}{}
	}

	return res
}

func perms(n, k int) int {
	return factorial(n) / factorial(n-k)
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}

	f, ok := factorials[n]
	if ok {
		return f
	}

	res := 1
	for i := n; i > 0; i-- {
		f, ok := factorials[i]
		if ok {
			res *= f
			break
		}
		res *= i
	}

	factorials[n] = res

	return res
}

func main() {
	fmt.Println(countSpecialNumbers(20))
}
