package remove_nth_from_end

import (
	"reflect"
	"testing"
)

func makeList(nums []int) *ListNode {
	if nums == nil {
		return nil
	}

	var res, tmp, prev *ListNode
	for _, num := range nums {
		tmp = &ListNode{
			Val: num,
		}

		if res == nil {
			res = tmp
		} else {
			prev.Next = tmp
		}
		prev = tmp
	}

	return res
}

func Test_removeNthFromEnd(t *testing.T) {
	type args struct {
		head *ListNode
		n    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "1",
			args: args{
				head: makeList([]int{1, 2, 3, 4, 5}),
				n:    2,
			},
			want: makeList([]int{1, 2, 3, 5}),
		},
		{
			name: "2",
			args: args{
				head: makeList([]int{1}),
				n:    1,
			},
			want: makeList([]int{}),
		},
		{
			name: "3",
			args: args{
				head: makeList([]int{1, 2}),
				n:    1,
			},
			want: makeList([]int{1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeNthFromEnd(tt.args.head, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeNthFromEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}
