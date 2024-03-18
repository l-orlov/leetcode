package pascal_triangle

import (
	"reflect"
	"testing"
)

func Test_generate(t *testing.T) {
	tests := []struct {
		name    string
		numRows int
		want    [][]int
	}{
		{
			name:    "1",
			numRows: 1,
			want:    [][]int{{1}},
		},
		{
			name:    "2",
			numRows: 2,
			want:    [][]int{{1}, {1, 1}},
		},
		{
			name:    "3",
			numRows: 3,
			want:    [][]int{{1}, {1, 1}, {1, 2, 1}},
		},
		{
			name:    "4",
			numRows: 4,
			want:    [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generate(tt.numRows); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generate() = %v, want %v", got, tt.want)
			}
		})
	}
}
