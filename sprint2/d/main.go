package main

type ListNode struct {
	data string
	next *ListNode
}

func Solution(head *ListNode, elem string) int {
	if head == nil {
		return -1
	}
	if head.data == elem {
		return 0
	}
	if head.next == nil {
		return -1
	}
	head = head.next
	idx := 1
	for head != nil {
		if head.data == elem {
			return idx
		}
		idx++
		head = head.next
	}
	return -1
}
