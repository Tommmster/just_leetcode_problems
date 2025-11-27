package main

// Given the head of a singly linked list, reverse the list, and return the reversed list.
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	next := head.Next
	head.Next = nil // The current head will be the last node of the reverse list

	return rev(next, head)
}

func rev(current *ListNode, previous *ListNode) *ListNode {
	if current.Next == nil {
		// Point backwards, return the new head
		current.Next = previous
		return current
	} else {
		p := rev(current.Next, current)
		current.Next = previous // Point backwards
		return p                // Return the new head
	}
}
