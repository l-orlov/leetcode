package main

import "fmt"

/*
pressure(1,0) -> pressure(0,0) / 2
pressure(1,1) -> pressure(0,0) / 2

pressure(2,0) -> (pressure(1,0) + pressure(0,0)) / 2
pressure(2,1) -> (pressure(1,0) + pressure(1,1)) / 2
pressure(2,2) -> (pressure(1,1)) / 2

pressure(3,0)
*/

var pressureMap = map[string]float64{}

func pressureMapKey(x, y int) string {
	return fmt.Sprintf("%d_%d", x, y)
}

func pressure(x, y int) float64 {
	if x == 0 && y == 0 {
		return 0
	}

	val, ok := pressureMap[pressureMapKey(x, y)]
	if ok {
		return val
	}

	var res float64
	if y == 0 {
		res = (1 + pressure(x-1, y)) / 2
	} else if y == x {
		res = (1 + pressure(x-1, y-1)) / 2
	} else {
		res = (2 + pressure(x-1, y) + pressure(x-1, y-1)) / 2
	}
	pressureMap[pressureMapKey(x, y)] = res

	return res
}
