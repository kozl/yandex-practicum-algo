package main

// <template>
type Node struct {
	value int
	left  *Node
	right *Node
}

// <template>

func insert(root *Node, key int) *Node {
	if root == nil {
		return &Node{value: key}
	}
	if key < root.value {
		root.left = insert(root.left, key)
	} else {
		root.right = insert(root.right, key)
	}
	return root
}
