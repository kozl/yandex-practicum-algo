package main

import (
	"fmt"
	"strings"
)

// <template>
type Node struct {
	value int
	left  *Node
	right *Node
}

// <template>

func printRange(root *Node, left int, right int) {
	arr := []int{}
	appendElementsFromRange(root, left, right, &arr)
	fmt.Println(format(arr))
}

func appendElementsFromRange(root *Node, left int, right int, arr *[]int) {
	if root.value >= left {
		if root.left != nil {
			appendElementsFromRange(root.left, left, right, arr)
		}
		if root.value <= right {
			*arr = append(*arr, root.value)
		}
	}
	if root.value <= right && root.right != nil {
		appendElementsFromRange(root.right, left, right, arr)
	}
}

func format(arr []int) string {
	tmp := make([]string, len(arr))
	for idx, el := range arr {
		tmp[idx] = fmt.Sprint(el)
	}
	return strings.Join(tmp, " ")
}
