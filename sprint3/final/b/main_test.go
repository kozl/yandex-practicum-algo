package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuickSort(t *testing.T) {
	array := []participant{
		{"alla", 4, 100},
		{"gena", 6, 1000},
		{"gosha", 2, 90},
		{"rita", 2, 90},
		{"timofey", 4, 80},
	}
	expected := []participant{
		{"rita", 2, 90},
		{"gosha", 2, 90},
		{"alla", 4, 100},
		{"timofey", 4, 80},
		{"gena", 6, 1000},
	}

	got := quicksort(array, lt)
	assert.Equal(t, expected, got)
}
