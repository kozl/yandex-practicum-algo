package main

// <template>
type Node struct {
	value int
	left  *Node
	right *Node
}

// <template>

func Solution(root1 *Node, root2 *Node) bool {
	return equal(root1, root2)
}

func equal(root1, root2 *Node) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 != nil && root2 != nil {
		if root1.value == root2.value {
			return equal(root1.left, root2.left) && equal(root1.right, root2.right)
		}
	}
	return false
}
