package main

import (
	"fmt"
)

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	if n == -1 {
		return 1 / x
	}

	val := myPow(x, n/2)
	res := val * val
	if n%2 != 0 {
		if n < 0 {
			res /= x
		} else {
			res *= x
		}
	}

	return res
}

func main() {
	fmt.Println(myPow(0, -3))
}
