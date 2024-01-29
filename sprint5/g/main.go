package main

// <template>
type Node struct {
	value int
	left  *Node
	right *Node
}

// <template>

func Solution(root *Node) int {
	panic("not implemented")
}

func score(root *Node, scrores *[]int) int {
	if root == nil {
		return 0
	}
	if root.left == nil && root.right == nil {
		if root.value <= 0 {
			return 0
		}
		return root.value
	}
	scoreLeft := score(root.left, scrores)
	scoreRight := score(root.right, scrores)
	*scrores = append(*scrores, scoreLeft+scoreRight)
	return max(scoreLeft, scoreRight)
}
