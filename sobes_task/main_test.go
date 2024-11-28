package main

import "testing"

func Test_pressure(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "0,0",
			args: args{
				x: 0,
				y: 0,
			},
			want: 0,
		},
		{
			name: "1,0",
			args: args{
				x: 1,
				y: 0,
			},
			want: 0.5,
		},
		{
			name: "1,1",
			args: args{
				x: 1,
				y: 1,
			},
			want: 0.5,
		},
		{
			name: "2,0",
			args: args{
				x: 2,
				y: 0,
			},
			want: 0.75,
		},
		{
			name: "2,1",
			args: args{
				x: 2,
				y: 1,
			},
			want: 1.5,
		},
		{
			name: "2,2",
			args: args{
				x: 2,
				y: 2,
			},
			want: 0.75,
		},
		{
			name: "3,0",
			args: args{
				x: 3,
				y: 0,
			},
			want: 0.875,
		},
		{
			name: "3,1",
			args: args{
				x: 3,
				y: 1,
			},
			want: 2.125,
		},
		{
			name: "3,2",
			args: args{
				x: 3,
				y: 2,
			},
			want: 2.125,
		},
		{
			name: "3,3",
			args: args{
				x: 3,
				y: 3,
			},
			want: 0.875,
		},
		{
			name: "322,156",
			args: args{
				x: 322,
				y: 156,
			},
			want: 306.48749781747574,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pressure(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("pressure() = %v, want %v", got, tt.want)
			}
		})
	}
}
