package remove_nth_from_end

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	firstNode := head
	secondNode := head
	i := 0
	for firstNode != nil {
		if i > n {
			secondNode = secondNode.Next
		}
		firstNode = firstNode.Next
		i += 1
	}

	if n == i {
		head = head.Next
	} else {
		secondNode.Next = secondNode.Next.Next
	}
	return head
}
