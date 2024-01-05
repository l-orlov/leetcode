package climb_stairs

import "testing"

func Test_climbStairs(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "1",
			n:    1,
			want: 1,
		},
		{
			name: "2",
			n:    2,
			want: 2,
		},
		{
			name: "3",
			n:    3,
			want: 3,
		},
		{
			name: "4",
			n:    4,
			want: 5,
		},
		{
			name: "5",
			n:    5,
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := climbStairs(tt.n); got != tt.want {
				t.Errorf("climbStairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
