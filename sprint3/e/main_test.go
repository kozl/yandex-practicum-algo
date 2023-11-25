package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve_Example1(t *testing.T) {
	houses := []int{999, 999, 999}
	budget := 300
	expect := 0
	got := solve(houses, budget)
	assert.Equal(t, expect, got)
}

func TestSolve_Example2(t *testing.T) {
	houses := []int{350, 999, 200}
	budget := 1000
	expect := 2
	got := solve(houses, budget)
	assert.Equal(t, expect, got)
}
