package main

import "testing"

func TestSolution_example1(t *testing.T) {
	node1 := Node{1, nil, nil}
	node2 := Node{-5, nil, nil}
	node3 := Node{3, &node1, &node2}
	node4 := Node{10, nil, nil}
	node5 := Node{2, &node3, &node4}
	if !Solution(&node5) {
		t.Fatal("WA")
	}
}

func TestSolution_example2(t *testing.T) {
	node8 := Node{8, nil, nil}
	node7 := Node{7, nil, nil}
	node6 := Node{6, nil, nil}
	node5 := Node{5, nil, nil}
	node4 := Node{4, &node7, &node8}
	node3 := Node{3, &node5, &node6}
	node2 := Node{2, nil, &node4}
	node1 := Node{1, &node3, nil}
	node0 := Node{0, &node1, &node2}
	if Solution(&node0) {
		t.Fatal("WA")
	}
}
