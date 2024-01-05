package main

import "fmt"

func main() {
	var (
		gas  []int
		cost []int
	)

	gas = []int{1, 2, 3, 4, 5}
	cost = []int{3, 4, 5, 1, 2}
	fmt.Println(canCompleteCircuit(gas, cost))

	//gas = []int{2, 3, 4}
	//cost = []int{3, 4, 3}
	//fmt.Println(canCompleteCircuit(gas, cost))
}

func canCompleteCircuit(gas []int, cost []int) int {
	var (
		diffByIdx       = make([]int, len(gas))
		startPoints     []int
		positiveDiffSum int
		negativeDiffSum int
		prevDiff        int
	)

	for idx := 0; idx < len(gas); idx++ {
		diff := gas[idx] - cost[idx]
		diffByIdx[idx] = diff
		if diff >= 0 && (idx == 0 || prevDiff < 0) {
			startPoints = append(startPoints, idx)
			positiveDiffSum += diff
		} else {
			negativeDiffSum += diff
		}
		prevDiff = diff
	}

	if positiveDiffSum < -negativeDiffSum {
		return -1
	}

	for _, idx := range startPoints {
		if canCompleteCircuitFromIdx(diffByIdx, idx) {
			return idx
		}
	}

	return -1
}

func canCompleteCircuitFromIdx(diffByIdx []int, idx int) bool {
	if idx >= len(diffByIdx) {
		return false
	}

	tank := 0
	for i := 0; i < len(diffByIdx); i++ {
		tank += diffByIdx[idx]
		if tank < 0 {
			return false
		}
		idx = (idx + 1) % len(diffByIdx)
	}

	return true
}
