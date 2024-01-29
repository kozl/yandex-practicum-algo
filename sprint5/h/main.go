package main

import (
	"fmt"
	"strconv"
)

// <template>
type Node struct {
	value int
	left  *Node
	right *Node
}

// <template>

func Solution(root *Node) int {
	nums := getNums(root, "", []string{})
	sum := 0
	for _, num := range nums {
		d, _ := strconv.Atoi(num)
		sum += d
	}
	return sum
}

func getNums(root *Node, s string, num []string) []string {
	if root == nil {
		return num
	}
	s += fmt.Sprint(root.value)
	if root.left == nil && root.right == nil {
		num = append(num, s)
		return num
	}
	num = getNums(root.left, s, num)
	num = getNums(root.right, s, num)
	return num
}
