package main

import (
	"bytes"
	"testing"
)

func TestSolve_example1(t *testing.T) {
	nodesCount := 6
	edges := []edge{
		{2, 6},
		{1, 6},
		{3, 1},
		{2, 5},
		{4, 3},
		{3, 2},
		{1, 2},
		{1, 4},
	}

	expected := `0 11
1 6
8 9
7 10
2 3
4 5
`

	var b bytes.Buffer
	solve(&b, nodesCount, edges)
	got := b.String()

	if got != expected {
		t.Fatalf("got = %q, expected = %q", got, expected)
	}
}

func TestSolve_example2(t *testing.T) {
	nodesCount := 3
	edges := []edge{
		{1, 2},
		{2, 3},
	}

	expected := `0 5
1 4
2 3
`

	var b bytes.Buffer
	solve(&b, nodesCount, edges)
	got := b.String()

	if got != expected {
		t.Fatalf("got = %q, expected = %q", got, expected)
	}
}
