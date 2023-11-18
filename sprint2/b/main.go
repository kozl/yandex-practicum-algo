package main

import "fmt"

type ListNode struct {
	data string
	next *ListNode
}

func Solution(head *ListNode) {
	fmt.Println(head.data)
	next := head.next
	for next != nil {
		fmt.Println(next.data)
		next = next.next
	}
}
