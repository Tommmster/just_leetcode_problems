package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func listFromSlice(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}
	head := &ListNode{Val: vals[0]}
	cur := head
	for _, v := range vals[1:] {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
	}
	return head
}

func sliceFromList(head *ListNode) []int {
	out := []int{}
	for head != nil {
		out = append(out, head.Val)
		head = head.Next
	}
	return out
}

// Find node at index i (0-based)
func nodeAt(head *ListNode, idx int) *ListNode {
	for idx > 0 && head != nil {
		head = head.Next
		idx--
	}
	return head
}

// Return the number of nodes starting from the head
func length(head *ListNode) int {
	// While a recursive implementation is much more elegant (Meine Meinung)
	// Go's compiler does not guarantee Tail Call Optimization, so we're
	// being safe and using an old fashioned for-loop
	var (
		count int = 0
		node      = head
	)

	for node != nil {
		count++
		node = node.Next
	}

	return count
}
