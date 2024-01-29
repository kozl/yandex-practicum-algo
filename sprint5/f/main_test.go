package main

import "testing"

func TestMain(t *testing.T) {
	node1 := Node{1, nil, nil}
	node2 := Node{4, nil, nil}
	node3 := Node{3, &node1, &node2}
	node4 := Node{8, nil, nil}
	node5 := Node{5, &node3, &node4}
	if Solution(&node5) != 3 {
		t.Fatal("WA")
	}
}
