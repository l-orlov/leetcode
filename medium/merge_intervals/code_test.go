package main

import "testing"

func Benchmark_mergeV1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mergeV1([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}})
	}
}

func Benchmark_mergeV2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mergeV2([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}})
	}
}
