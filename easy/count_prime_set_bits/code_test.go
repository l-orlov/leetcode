package count_prime_set_bits

import "testing"

func Test_countPrimeSetBits(t *testing.T) {
	tests := []struct {
		name  string
		left  int
		right int
		want  int
	}{
		{
			name:  "1",
			left:  6,
			right: 10,
			want:  4,
		},
		{
			name:  "2",
			left:  10,
			right: 15,
			want:  5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countPrimeSetBits(tt.left, tt.right); got != tt.want {
				t.Errorf("countPrimeSetBits() = %v, want %v", got, tt.want)
			}
		})
	}
}
