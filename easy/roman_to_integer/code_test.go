package roman_to_integer

import "testing"

func Test_romanToInt(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			s:    "I",
			want: 1,
		},
		{
			s:    "V",
			want: 5,
		},
		{
			s:    "X",
			want: 10,
		},
		{
			s:    "L",
			want: 50,
		},
		{
			s:    "C",
			want: 100,
		},
		{
			s:    "D",
			want: 500,
		},
		{
			s:    "II",
			want: 2,
		},
		{
			s:    "XII",
			want: 12,
		},
		{
			s:    "XXVII",
			want: 27,
		},
		{
			s:    "LVIII",
			want: 58,
		},
		{
			s:    "MCMXCIV",
			want: 1994,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := romanToInt(tt.s); got != tt.want {
				t.Errorf("romanToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
