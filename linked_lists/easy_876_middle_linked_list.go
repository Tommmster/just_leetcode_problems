package main

func middleNode(head *ListNode) *ListNode {
	l := length(head)
	if l == 1 {
		return head
	} else if l == 2 {
		return head.Next
	}

	var want = l / 2
	for i := 0; i < want; i++ {
		head = head.Next
	}

	return head
}
