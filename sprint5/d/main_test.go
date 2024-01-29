package main

import "testing"

func TestMain(t *testing.T) {
	node1 := Node{1, nil, nil}
	node2 := Node{2, nil, nil}
	node3 := Node{3, &node1, &node2}
	node4 := Node{1, nil, nil}
	node5 := Node{2, nil, nil}
	node6 := Node{3, &node4, &node5}

	if !Solution(&node3, &node6) {
		t.Fatal("WA")
	}
}
