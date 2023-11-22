package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCustomSortWithPrint_Case1(t *testing.T) {
	input := []int{15, 56, 2}
	w := bytes.Buffer{}
	customSort(input)
	printArray(&w, input)
	assert.Equal(t, "56215\n", w.String())
}

func TestCustomSortWithPrint_Case2(t *testing.T) {
	input := []int{1, 783, 2}
	w := bytes.Buffer{}
	customSort(input)
	printArray(&w, input)
	assert.Equal(t, "78321\n", w.String())
}

func TestCustomSortWithPrint_Case3(t *testing.T) {
	input := []int{2, 4, 5, 2, 10}
	w := bytes.Buffer{}
	customSort(input)
	printArray(&w, input)
	assert.Equal(t, "542210\n", w.String())
}

func TestCustomSortWithPrint_Case4(t *testing.T) {
	input := []int{9, 10, 1, 1, 1, 6}
	w := bytes.Buffer{}
	customSort(input)
	printArray(&w, input)
	assert.Equal(t, "9611110\n", w.String())
}

func TestCustomSortWithPrint_Case5(t *testing.T) {
	input := []int{9, 6, 43, 81, 66, 69, 15, 33, 6, 53, 93, 64, 33, 88, 39, 34, 57, 23, 42, 44, 79, 25}
	w := bytes.Buffer{}
	customSort(input)
	printArray(&w, input)
	assert.Equal(t, "99388817969666664575344434239343333252315\n", w.String())
}
