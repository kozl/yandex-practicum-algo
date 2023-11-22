package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSort_case1(t *testing.T) {
	w := bytes.Buffer{}
	expected := `3 4 2 1 9
3 2 1 4 9
2 1 3 4 9
1 2 3 4 9
`
	array := []int{4, 3, 9, 2, 1}
	bubbleSortAndPrint(&w, array)
	assert.Equal(t, expected, w.String())
}

func TestBubbleSort_case2(t *testing.T) {
	w := bytes.Buffer{}
	expected := `8 9 10 11 12
`
	array := []int{12, 8, 9, 10, 11}
	bubbleSortAndPrint(&w, array)
	assert.Equal(t, expected, w.String())
}

func TestBubbleSort_case3(t *testing.T) {
	w := bytes.Buffer{}
	expected := `4 5
`
	array := []int{4, 5}
	bubbleSortAndPrint(&w, array)
	assert.Equal(t, expected, w.String())
}

func TestBubbleSort_case4(t *testing.T) {
	w := bytes.Buffer{}
	expected := `1 1 1 1 1
`
	array := []int{1, 1, 1, 1, 1}
	bubbleSortAndPrint(&w, array)
	assert.Equal(t, expected, w.String())
}
