package main

import (
	"bytes"
	"testing"
)

func TestPrintAdjacencyMatrix(t *testing.T) {
	nodesCount := 5
	edges := []edge{
		{1, 3},
		{2, 3},
		{5, 2},
	}
	expected := `0 0 1 0 0
0 0 1 0 0
0 0 0 0 0
0 0 0 0 0
0 1 0 0 0
`
	var b bytes.Buffer
	printAdjacencyMatrix(&b, nodesCount, edges)
	got := b.String()

	if got != expected {
		t.Fatalf("got = %q, expected = %q", got, expected)
	}
}
