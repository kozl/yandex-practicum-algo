package main

// <templatet>
type Node struct {
	value int
	left  *Node
	right *Node
}

// <template>

func remove(node *Node, key int) *Node {
	if node == nil {
		return node
	}
	if key == node.value {
		if node.left == nil && node.right == nil {
			return nil
		}
		if node.left == nil {
			return node.right
		}
		if node.right == nil {
			return node.left
		}
		var val int
		node.right, val = deleteLeftMostNode(node.right)
		node.value = val
	} else if key > node.value {
		node.right = remove(node.right, key)
	} else {
		node.left = remove(node.left, key)
	}
	return node
}

func deleteLeftMostNode(node *Node) (*Node, int) {
	if node.left == nil {
		if node.right == nil {
			return nil, node.value
		}
		return node.right, node.value
	}
	var val int
	node.left, val = deleteLeftMostNode(node.left)
	return node, val
}
