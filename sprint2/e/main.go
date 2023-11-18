package main

type ListNode struct {
	data string
	next *ListNode
	prev *ListNode
}

func Solution(head *ListNode) *ListNode {
	previous := head.prev
	for head.next != nil {
		head.next, head.prev = previous, head.next
		previous = head
		head = head.prev
	}
	head.prev = nil
	head.next = previous

	return head
}
