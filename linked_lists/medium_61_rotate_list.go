package main

// Given the head of a linked list, rotate the list to the right by k places.
func rotateRight(head *ListNode, k int) *ListNode {

	// k indicates that the last k nodes should be moved up front. For example, if k = 1
	// [1,2,3,4,5] -> // [5] -> [1] -> [2]-> [3] -> [4]
	llen := 0
	for t := head; t != nil; t = t.Next {
		llen++
	}
	k = k % llen

	if head == nil || head.Next == nil || k == 0 {
		return head
	}
	// sublist will return the new head of the list
	nh, nl, _ := sublist(head, k)
	nl.Next = head
	return nh
}

func sublist(current *ListNode, k int) (*ListNode, *ListNode, int) {
	if current.Next == nil {
		return nil, current, k - 1
	}
	// If the current node is not the last, then keep nh and nl, candidates for head and last element
	// of the sublist
	nh, nl, count := sublist(current.Next /* not nil */, k)
	if count > 0 {
		return nil, nl, count - 1
	} else {
		if nh != nil {
			return nh, nl, 0
		}

		nh = current.Next
		current.Next = nil
		return nh, nl, 0
	}
}
