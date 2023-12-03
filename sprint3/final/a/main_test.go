package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearch(t *testing.T) {
	array := []int{1, 2, 4, 6, 9, 10, 15, 20}
	n := 9
	got := binarySearch(array, 0, len(array)-1, func(array []int, i int) int {
		if array[i] == n {
			return 0
		}
		if array[i] < n {
			return 1
		}
		return -1
	})
	assert.Equal(t, 4, got)
}

func TestBrokenSearch_example1(t *testing.T) {
	array := []int{9, 10, 15, 20, 1, 2, 4, 6}
	n := 2
	got := brokenSearch(array, n)
	assert.Equal(t, 5, got)
}

func TestBrokenSearch_example2(t *testing.T) {
	array := []int{9, 10, 15, 20, 21, 22, 24, 6}
	n := 6
	got := brokenSearch(array, n)
	assert.Equal(t, 7, got)
}

func TestBrokenSearch_example3(t *testing.T) {
	array := []int{3, 1}
	n := 1
	got := brokenSearch(array, n)
	assert.Equal(t, 1, got)
}

func TestBrokenSearch_example4(t *testing.T) {
	array := []int{3}
	n := 3
	got := brokenSearch(array, n)
	assert.Equal(t, 0, got)
}

func TestBrokenSearch_example5(t *testing.T) {
	array := []int{1, 2, 3, 4, 5, 6}
	n := 3
	got := brokenSearch(array, n)
	assert.Equal(t, 2, got)
}
func TestBrokenSearch_example6(t *testing.T) {
	array := []int{3, 5, 6, 7, 9, 1, 2}
	n := 4
	got := brokenSearch(array, n)
	assert.Equal(t, -1, got)
}
