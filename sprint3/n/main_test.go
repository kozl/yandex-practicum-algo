package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortAndMergeAreas_Case1(t *testing.T) {
	input := []area{
		{7, 8},
		{7, 8},
		{2, 3},
		{6, 10},
	}
	expected := []area{
		{2, 3},
		{6, 10},
	}

	sortAreas(input)
	got := mergeAreas(input)
	assert.Equal(t, expected, got)
}

func TestSortAndMergeAreas_Case2(t *testing.T) {
	input := []area{
		{2, 3},
		{5, 6},
		{3, 4},
		{3, 4},
	}
	expected := []area{
		{2, 4},
		{5, 6},
	}

	sortAreas(input)
	got := mergeAreas(input)
	assert.Equal(t, expected, got)
}

func TestSortAndMergeAreas_Case3(t *testing.T) {
	input := []area{
		{1, 3},
		{3, 5},
		{4, 6},
		{5, 6},
		{2, 4},
		{7, 10},
	}
	expected := []area{
		{1, 6},
		{7, 10},
	}

	sortAreas(input)
	got := mergeAreas(input)
	assert.Equal(t, expected, got)
}
