package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// Given the head of a singly linked list, reverse the list, and return the reversed list.
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	return rev(head.Next, head)
}

func rev(current *ListNode, prev *ListNode) *ListNode {
	if current.Next == nil {
		current.Next = prev
		return current
	} else {
		p := rev(current.Next, current)
		current.Next = prev
		return p
	}
}
