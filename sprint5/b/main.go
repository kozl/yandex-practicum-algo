package main

import "math"

// <template>
type Node struct {
	value int
	left  *Node
	right *Node
}

// <template>

func Solution(root *Node) bool {
	return isBalanced(root)
}

func isBalanced(root *Node) bool {
	if root == nil {
		return true
	}
	if !isBalanced(root.left) {
		return false
	}
	if int(math.Abs(float64(height(root.left)-height(root.right)))) > 1 {
		return false
	}
	if !isBalanced(root.right) {
		return false
	}
	return true
}

func height(root *Node) int {
	if root == nil {
		return 0
	}
	return max(height(root.left), height(root.right)) + 1
}
