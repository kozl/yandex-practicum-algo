package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve_Example1(t *testing.T) {
	children := []int{1, 2}
	cookies := []int{2, 1, 3}
	expect := 2
	got := solve(children, cookies)
	assert.Equal(t, expect, got)
}

func TestSolve_Example2(t *testing.T) {
	children := []int{2, 1, 3}
	cookies := []int{1, 1}
	expect := 1
	got := solve(children, cookies)
	assert.Equal(t, expect, got)
}

func TestSolve_Example3(t *testing.T) {
	children := []int{4}
	cookies := []int{1, 4, 7, 10, 2, 2, 7, 8}
	expect := 1
	got := solve(children, cookies)
	assert.Equal(t, expect, got)
}
