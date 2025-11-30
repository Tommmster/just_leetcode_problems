package main

// Given the head of a linked list, return the list after sorting it in ascending order.
func SortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	l := length(head)

	return doSort(head, l)
}

// Sort the elements of the list pointed by head, return the head of the sorted list
func doSort(head *ListNode, l /* size */ int) *ListNode {
	if l <= 1 {
		return head
	} else if next := head.Next; l == 2 && next.Val < head.Val {
		head.Val, next.Val = next.Val, head.Val
	}

	low, high := split(head, l/2)
	low = doSort(low, l/2)
	high = doSort(high, l-l/2)

	return merge(low, high)
}

func merge(low, high *ListNode) *ListNode {
	var head *ListNode
	if low.Val <= high.Val {
		head = low
		low = low.Next
	} else {
		head = high
		high = high.Next
	}

	t := head
	for low != nil && high != nil {
		if low.Val < high.Val {
			t.Next = low
			low = low.Next
		} else {
			t.Next = high
			high = high.Next
		}

		t = t.Next
	}
	if low != nil {
		t.Next = low
	} else {
		t.Next = high
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

// split the linked list pointed by head in two lists, the first one of length l
func split(head *ListNode, l int) (*ListNode, *ListNode) {
	var (
		hl = head
		hh = head
	)

	// quick wins
	if l == 0 {
		return nil, head
	} else if head == nil {
		return nil, nil
	}

	for i := 0; i < l; i++ {
		hh = hh.Next
		if hh == nil {
			return hl, nil
		}
	}

	{
		t := hl
		for t.Next != hh {
			t = t.Next
		}
		t.Next = nil
	}

	return hl, hh
}
