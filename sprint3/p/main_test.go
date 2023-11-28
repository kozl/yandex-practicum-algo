package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution_Example1(t *testing.T) {
	array := []int{0, 1, 3, 2}
	expect := 3
	got := solution(array)
	assert.Equal(t, expect, got)
}
func TestSolution_Example2(t *testing.T) {
	array := []int{3, 6, 7, 4, 1, 5, 0, 2}
	expect := 1
	got := solution(array)
	assert.Equal(t, expect, got)
}

func TestSolution_Example3(t *testing.T) {
	array := []int{1, 0, 2, 3, 4}
	expect := 4
	got := solution(array)
	assert.Equal(t, expect, got)
}
