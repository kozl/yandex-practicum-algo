package main

import (
	"bytes"
	"testing"
)

func TestPrintAdjacencyList(t *testing.T) {
	nodesCount := 5
	edges := []edge{
		{1, 3},
		{2, 3},
		{5, 2},
	}
	expected := `1 3
1 3
0
0
1 2
`

	var b bytes.Buffer
	printAdjacencyList(&b, nodesCount, edges)
	got := b.String()

	if got != expected {
		t.Fatalf("got = %q, expected = %q", got, expected)
	}
}
