package candy

import "testing"

func Test_candy(t *testing.T) {
	tests := []struct {
		name    string
		ratings []int
		want    int
	}{
		{
			name:    "1",
			ratings: []int{60, 80, 100, 100, 100, 100, 100},
			// 1, 2, 3, 1, 1, 1, 1
			want: 10,
		},
		{
			name:    "2",
			ratings: []int{1, 1, 1, 1},
			// 1, 1, 1, 1
			want: 4,
		},
		{
			name:    "3",
			ratings: []int{1, 0, 2},
			// 2, 1, 2
			want: 5,
		},
		{
			name:    "4",
			ratings: []int{1, 2, 2},
			// 1, 2, 1
			want: 4,
		},
		{
			name:    "5",
			ratings: []int{1, 0, 2, 1, 3, 4, 5, 5, 4, 4, 1, 2},
			// 2, 1, 2, 1, 2, 3, 4, 2, 1, 2, 1, 2
			want: 23,
		},
		{
			name:    "6",
			ratings: []int{1, 3, 4, 5, 2},
			// 1, 2, 3, 4, 1
			want: 11,
		},
		{
			name:    "7",
			ratings: []int{1, 6, 10, 8, 7, 3, 2},
			// 1, 2, 5, 4, 3, 2, 1
			want: 18,
		},
		{
			name:    "8",
			ratings: []int{1, 2, 87, 87, 87, 2, 1},
			// 1, 2, 3, 1, 3, 2, 1
			want: 13,
		},
		{
			name:    "9",
			ratings: []int{1, 2, 3, 5, 4, 3, 2, 1},
			// 1, 2, 3, 5, 4, 3, 2, 1
			want: 21,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := candy(tt.ratings); got != tt.want {
				t.Errorf("candy() = %v, want %v", got, tt.want)
			}
		})
	}
}
