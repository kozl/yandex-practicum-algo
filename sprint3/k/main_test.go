package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMerge(t *testing.T) {
	array := []int{1, 3, 5, 6, 2, 4, 8, 10, 11}
	got := merge(array, 0, 4, 9)
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 8, 10, 11}, got)

	array = []int{1, 4, 9, 2, 10, 11}
	got = merge(array, 0, 3, 6)
	assert.Equal(t, []int{1, 2, 4, 9, 10, 11}, got)
}

func TestMergeSort(t *testing.T) {
	array := []int{1, 4, 2, 10, 1, 2}
	merge_sort(array, 0, 6)
	assert.Equal(t, []int{1, 1, 2, 2, 4, 10}, array)
}
