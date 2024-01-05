package main

import (
	"fmt"
)

func countGoodNumbers(n int64) int {
	var modulo int64 = 1_000_000_007

	return int(pow(5, (n+1)/2, modulo) * pow(4, n/2, modulo) % modulo)
}

func pow(a, b, modulo int64) int64 {
	if b == 0 {
		return 1
	}
	if b == 1 {
		return a
	}

	x := pow(a, b/2, modulo)
	if b%2 == 0 {
		return x * x % modulo
	}

	return (a * x % modulo) * x % modulo
}

func main() {
	fmt.Println(countGoodNumbers(3))
}
