package merge_two_sorted_lists

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

func Test_mergeTwoLists(t *testing.T) {
	tests := []struct {
		name  string
		list1 *ListNode
		list2 *ListNode
		want  *ListNode
	}{
		{
			name:  "1",
			list1: makeList([]int{1, 2, 4}),
			list2: makeList([]int{1, 3, 4}),
			want:  makeList([]int{1, 1, 2, 3, 4, 4}),
		},
		{
			name:  "2",
			list1: makeList([]int{}),
			list2: makeList([]int{}),
			want:  makeList([]int{}),
		},
		{
			name:  "3",
			list1: makeList([]int{}),
			list2: makeList([]int{0}),
			want:  makeList([]int{0}),
		},
		{
			name:  "4",
			list1: makeList([]int{0}),
			list2: makeList([]int{}),
			want:  makeList([]int{0}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTwoLists(tt.list1, tt.list2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTwoLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
