package main

type ListNode struct {
	data string
	next *ListNode
}

func Solution(head *ListNode, idx int) *ListNode {
	if idx == 0 {
		head = head.next
		return head
	}
	previous := getNodeByIndex(head, idx-1)
	if previous == nil || previous.next == nil {
		return head
	}
	previous.next = previous.next.next
	return head
}

func getNodeByIndex(node *ListNode, idx int) *ListNode {
	for idx > 0 {
		node = node.next
		if node == nil {
			break
		}
		idx--
	}
	return node
}
