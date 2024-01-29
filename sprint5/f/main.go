package main

// <template>
type Node struct {
	value int
	left  *Node
	right *Node
}

// <template>

func Solution(root *Node) int {
	return depth(root)
}

func depth(root *Node) int {
	if root == nil {
		return 0
	}
	return max(depth(root.left), depth(root.right)) + 1
}
