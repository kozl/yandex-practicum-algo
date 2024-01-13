package main

import "testing"

func TestRemove_example1(t *testing.T) {
	node1 := Node{2, nil, nil}
	node2 := Node{3, &node1, nil}
	node3 := Node{1, nil, &node2}
	node4 := Node{6, nil, nil}
	node5 := Node{8, &node4, nil}
	node6 := Node{10, &node5, nil}
	node7 := Node{5, &node3, &node6}
	newHead := remove(&node7, 10)
	if newHead.value != 5 {
		t.Error("WA")
	}
	if newHead.right != &node5 {
		t.Error("WA")
	}
	if newHead.right.value != 8 {
		t.Error("WA")
	}
}

func TestRemove_example2(t *testing.T) {
	node10 := Node{99, nil, nil}
	node9 := Node{72, nil, nil}
	node8 := Node{91, &node9, &node10}
	node7 := Node{50, nil, nil}
	node6 := Node{32, nil, nil}
	node5 := Node{29, nil, &node6}
	node4 := Node{11, nil, nil}
	node3 := Node{65, &node7, &node8}
	node2 := Node{20, &node4, &node5}
	node1 := Node{41, &node2, &node3}
	newHead := remove(&node1, 41)
	if newHead.value != 50 {
		t.Error("WA")
	}
	if node3.left != nil {
		t.Error("WA")
	}
}

func TestRemove_example3(t *testing.T) {
	node10 := Node{932, nil, nil}
	node9 := Node{912, nil, &node10}
	node8 := Node{822, nil, nil}
	node7 := Node{870, &node8, &node9}
	node6 := Node{701, nil, nil}
	node5 := Node{702, &node6, &node7}
	node4 := Node{266, nil, nil}
	node3 := Node{191, nil, &node4}
	node2 := Node{298, &node3, nil}
	node1 := Node{668, &node2, &node5}
	newHead := remove(&node1, 545)
	if newHead.value != 668 {
		t.Error("WA")
	}
}
