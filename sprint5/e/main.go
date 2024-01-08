package main

// <template>
type Node struct {
	value int
	left  *Node
	right *Node
}

// <template>

func Solution(root *Node) bool {
	_, ok := isSearchTree(root)
	return ok
}

func isSearchTree(root *Node) ([]Node, bool) {
	if root == nil {
		return []Node{}, true
	}
	if root.left == nil && root.right == nil {
		return []Node{*root}, true
	}

	leftSubtree, ok := isSearchTree(root.left)
	if !ok {
		return nil, false
	}

	rightSubtree, ok := isSearchTree(root.right)
	if !ok {
		return nil, false
	}

	for _, node := range leftSubtree {
		if node.value >= root.value {
			return nil, false
		}
	}

	for _, node := range rightSubtree {
		if node.value <= root.value {
			return nil, false
		}
	}

	return append(leftSubtree, rightSubtree...), true
}
