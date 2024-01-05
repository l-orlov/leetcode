package binary_search

import "testing"

func Test_search(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				nums:   nil,
				target: 9,
			},
			want: -1,
		},
		{
			args: args{
				nums:   []int{},
				target: 9,
			},
			want: -1,
		},
		{
			args: args{
				nums:   []int{9},
				target: 9,
			},
			want: 0,
		},
		{
			args: args{
				nums:   []int{0},
				target: 9,
			},
			want: -1,
		},
		{
			args: args{
				nums:   []int{-1, 0, 3, 5, 9, 12},
				target: 9,
			},
			want: 4,
		},
		{
			args: args{
				nums:   []int{-1, 0, 3, 5, 9, 12},
				target: 2,
			},
			want: -1,
		},
		{
			name: "",
			args: args{
				nums:   []int{2, 5},
				target: 2,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}
