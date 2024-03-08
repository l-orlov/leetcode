package max_product_diff

import "testing"

func Test_maxProductDifference(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "1",
			nums: []int{5, 6, 2, 7, 4},
			want: 34,
		},
		{
			name: "2",
			nums: []int{4, 2, 5, 9, 7, 4, 8},
			want: 64,
		},
		{
			name: "3",
			nums: []int{5, 9, 4, 6},
			want: 34,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProductDifference(tt.nums); got != tt.want {
				t.Errorf("maxProductDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
