package main

type Node struct {
	value int
	left  *Node
	right *Node
}

func Solution(root *Node) int {
	return findMax(root)
}

func findMax(root *Node) int {
	if root.left != nil && root.right != nil {
		return max(findMax(root.left), findMax(root.right), root.value)
	}
	if root.left != nil && root.right == nil {
		return max(findMax(root.left), root.value)
	}
	if root.right != nil && root.left == nil {
		return max(findMax(root.right), root.value)
	}
	return root.value
}
