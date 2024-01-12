package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_appendElementsFromRange(t *testing.T) {
	node1 := Node{2, nil, nil}
	node2 := Node{1, nil, &node1}
	node3 := Node{8, nil, nil}
	node4 := Node{8, nil, &node3}
	node5 := Node{9, &node4, nil}
	node6 := Node{10, &node5, nil}
	node7 := Node{5, &node2, &node6}
	arr := []int{}
	appendElementsFromRange(&node7, 2, 8, &arr)
	assert.Equal(t, []int{2, 5, 8, 8}, arr)
}
