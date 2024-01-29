package main

// <template>
type Node struct {
	value int
	left  *Node
	right *Node
}

// <template>

func Solution(root *Node) bool {
	if root == nil {
		return true
	}
	return simmetric(root.left, root.right)
}

func simmetric(root1, root2 *Node) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 != nil && root2 != nil {
		if root1.value == root2.value {
			return simmetric(root1.left, root2.right) && simmetric(root2.left, root1.right)
		}
	}
	return false
}
