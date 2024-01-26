package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSort_example1(t *testing.T) {
	arr := []person{
		{"alla", 4, 100},
		{"gena", 6, 1000},
		{"gosha", 2, 90},
		{"rita", 2, 90},
		{"timofey", 4, 80},
	}

	expected := []person{
		{"gena", 6, 1000},
		{"timofey", 4, 80},
		{"alla", 4, 100},
		{"gosha", 2, 90},
		{"rita", 2, 90},
	}

	got := sort(arr)
	assert.Equal(t, expected, got)
}

func TestSort_example2(t *testing.T) {
	arr := []person{
		{"alla", 0, 0},
		{"gena", 0, 0},
		{"gosha", 0, 0},
		{"rita", 0, 0},
		{"timofey", 0, 0},
	}

	expected := []person{
		{"alla", 0, 0},
		{"gena", 0, 0},
		{"gosha", 0, 0},
		{"rita", 0, 0},
		{"timofey", 0, 0},
	}

	got := sort(arr)
	assert.Equal(t, expected, got)
}
