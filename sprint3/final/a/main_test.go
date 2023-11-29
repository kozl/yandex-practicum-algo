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

func TestSolution_example1(t *testing.T) {
	array := []int{9, 10, 15, 20, 1, 2, 4, 6}
	n := 2
	got := solution(array, n)
	assert.Equal(t, 5, got)
}

func TestSolution_example2(t *testing.T) {
	array := []int{9, 10, 15, 20, 21, 22, 24, 6}
	n := 6
	got := solution(array, n)
	assert.Equal(t, 7, got)
}

func TestSolution_example3(t *testing.T) {
	array := []int{3, 1}
	n := 1
	got := solution(array, n)
	assert.Equal(t, 1, got)
}

func TestSolution_example4(t *testing.T) {
	array := []int{3}
	n := 3
	got := solution(array, n)
	assert.Equal(t, 0, got)
}

func TestSolution_example5(t *testing.T) {
	array := []int{1, 2, 3, 4, 5, 6}
	n := 3
	got := solution(array, n)
	assert.Equal(t, 2, got)
}
