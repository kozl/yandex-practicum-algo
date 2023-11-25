package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve_Example1(t *testing.T) {
	array := []int{1, 2, 3, 1, 2, 3, 4}
	n := 3
	expect := []int{1, 2, 3}
	got := getTopN(array, n)
	assert.Equal(t, expect, got)
}
func TestSolve_Example2(t *testing.T) {
	array := []int{1, 1, 1, 2, 2, 3}
	n := 1
	expect := []int{1}
	got := getTopN(array, n)
	assert.Equal(t, expect, got)
}

func TestSolve_Example3(t *testing.T) {
	array := []int{1, 2, 3}
	n := 3
	expect := []int{1, 2, 3}
	got := getTopN(array, n)
	assert.Equal(t, expect, got)
}
