package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountSort1(t *testing.T) {
	in := []int{0, 2, 1, 2, 0, 0, 1}
	expected := []int{0, 0, 0, 1, 1, 2, 2}
	countSort(in)
	assert.Equal(t, expected, in)
}

func TestCountSort2(t *testing.T) {
	in := []int{2, 1, 2, 0, 1}
	expected := []int{0, 1, 1, 2, 2}
	countSort(in)
	assert.Equal(t, expected, in)
}

func TestCountSort3(t *testing.T) {
	in := []int{2, 1, 1, 2, 0, 2}
	expected := []int{0, 1, 1, 2, 2, 2}
	countSort(in)
	assert.Equal(t, expected, in)
}

func TestCountSort4(t *testing.T) {
	in := []int{}
	expected := []int{}
	countSort(in)
	assert.Equal(t, expected, in)
}