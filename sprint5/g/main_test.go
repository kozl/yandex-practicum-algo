package main

import "testing"

func TestMain(t *testing.T) {
	node1 := Node{5, nil, nil}
	node2 := Node{1, nil, nil}
	node3 := Node{-3, &node2, &node1}
	node4 := Node{2, nil, nil}
	node5 := Node{2, &node4, &node3}
	if Solution(&node5) != 6 {
		t.Fatal("WA")
	}
}
