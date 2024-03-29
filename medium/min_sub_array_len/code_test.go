package min_sub_array_len

import "testing"

func Test_minSubArrayLen(t *testing.T) {
	type args struct {
		target int
		nums   []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "0",
			args: args{
				target: 15,
				nums:   []int{1, 2, 3, 5, 10, 2, 4, 8, 7, 5, 3, 2, 4, 10, 2},
			},
			want: 2,
		},
		{
			name: "1",
			args: args{
				target: 7,
				nums:   []int{2, 3, 1, 2, 4, 3},
			},
			want: 2,
		},
		{
			name: "2",
			args: args{
				target: 4,
				nums:   []int{1, 4, 4},
			},
			want: 1,
		},
		{
			name: "3",
			args: args{
				target: 11,
				nums:   []int{1, 1, 1, 1, 1, 1, 1, 1},
			},
			want: 0,
		},
		{
			name: "3",
			args: args{
				target: 213,
				nums:   []int{12, 28, 83, 4, 25, 26, 25, 2, 25, 25, 25, 12},
				// nums:   []int{12,28,83,4,25,26,25,2,25,25,25,12},
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSubArrayLen(tt.args.target, tt.args.nums); got != tt.want {
				t.Errorf("minSubArrayLen() = %v, want %v", got, tt.want)
			}
		})
	}
}
