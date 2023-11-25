package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve_Example1(t *testing.T) {
	array := []int{6, 3, 3, 2}
	expect := 8
	got := solve(array)
	assert.Equal(t, expect, got)
}

func TestSolve_Example2(t *testing.T) {
	array := []int{5, 3, 7, 2, 8, 3}
	expect := 20
	got := solve(array)
	assert.Equal(t, expect, got)
}
