package main

import "testing"

func Test_insert(t *testing.T) {
	node1 := Node{7, nil, nil}
	node2 := Node{8, &node1, nil}
	node3 := Node{7, nil, &node2}
	newHead := insert(&node3, 6)
	if newHead != &node3 {
		t.Error("WA")
	}
	if newHead.left.value != 6 {
		t.Error("WA")
	}
}
