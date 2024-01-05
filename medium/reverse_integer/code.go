package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	res := 0
	for x != 0 {
		tmpNum := x % 10
		tmpRes := res * 10

		// check overlapping
		if res > math.MaxInt32/10 || res < math.MinInt32/10 ||
			tmpRes > math.MaxInt32-tmpNum || tmpRes < math.MinInt32-tmpNum {
			return 0
		}

		res = tmpRes + tmpNum
		x /= 10
	}

	return res
}

func reverse32(x int32) int32 {
	//math.MinInt32 = -2147483648
	//math.MaxInt32 = 2147483647
	var res int32 = 0
	for x != 0 {
		tmp := x % 10
		if res > math.MaxInt32/10 || res < math.MinInt32/10 ||
			res*10 > math.MaxInt32-tmp || res*10 < math.MinInt32-tmp {
			return 0
		}
		res = res*10 + x%10
		x /= 10
	}

	return res
}

func main() {
	fmt.Println(reverse32(1000000003))
	// fmt.Println(reverse(-1234))
}
