package peak_mountain_array

import "testing"

func Test_peakOfMountainArray(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want int
	}{
		{
			name: "1",
			arr:  []int{0, 1, 2, 3, 2, 1, 0},
			want: 3,
		},
		{
			name: "2",
			arr:  []int{0, 10, 3, 2, 1, 0},
			want: 1,
		},
		{
			name: "3",
			arr:  []int{0, 1, 2, 12, 22, 32, 42, 52, 62, 72, 82, 92, 102, 112, 122, 132, 133, 132, 111, 0},
			want: 16,
		},
		{
			name: "4",
			arr:  []int{0, 1, 2, 3, 2, 1, 0},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := peakOfMountainArray(tt.arr); got != tt.want {
				t.Errorf("peakOfMountainArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
