package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve_Example1(t *testing.T) {
	array := []int{2, 3, 4}
	n := 2
	expect := 1
	got := getN(array, n)
	assert.Equal(t, expect, got)
}
func TestSolve_Example2(t *testing.T) {
	array := []int{1, 3, 1}
	n := 1
	expect := 0
	got := getN(array, n)
	assert.Equal(t, expect, got)
}

func TestSolve_Example3(t *testing.T) {
	array := []int{1, 3, 5}
	n := 3
	expect := 4
	got := getN(array, n)
	assert.Equal(t, expect, got)
}

func TestSolve_Example4(t *testing.T) {
	array := []int{7, 3, 8, 9, 10}
	n := 3
	expect := 1
	got := getN(array, n)
	assert.Equal(t, expect, got)
}

func TestSolve_Example5(t *testing.T) {
	array := []int{7, 2, 7, 3}
	n := 4
	expect := 4
	got := getN(array, n)
	assert.Equal(t, expect, got)
}

func TestSolve_Example6(t *testing.T) {
	array := []int{21, 91, 8, 68, 35, 71, 32, 49, 6}
	n := 22
	expect := 39
	got := getN(array, n)
	assert.Equal(t, expect, got)
}

func TestSolve_Example7(t *testing.T) {
	array := []int{27,67,3,4,17,98}
	n := 7
	expect := 31
	got := getN(array, n)
	assert.Equal(t, expect, got)
}