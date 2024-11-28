package newspapers_split

import "testing"

func Test_newspapersSplit(t *testing.T) {
	type args struct {
		newspapersReadTimes []int
		numCoworkers        int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				newspapersReadTimes: []int{7, 2, 5, 10, 8},
				numCoworkers:        2,
			},
			want: 18,
		},
		{
			name: "2",
			args: args{
				newspapersReadTimes: []int{2, 3, 5, 7},
				numCoworkers:        3,
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newspapersSplit(tt.args.newspapersReadTimes, tt.args.numCoworkers); got != tt.want {
				t.Errorf("newspapersSplit() = %v, want %v", got, tt.want)
			}
		})
	}
}
