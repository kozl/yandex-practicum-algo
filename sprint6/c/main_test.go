package main

import (
	"bytes"
	"testing"
)

func TestDFS_example1(t *testing.T) {
	nodesCount := 4
	edges := []edge{
		{3, 2},
		{4, 3},
		{1, 4},
		{1, 2},
	}
	startNode := 3

	expected := `3 2 1 4`

	var b bytes.Buffer
	dfs(&b, nodesCount, startNode, edges)
	got := b.String()

	if got != expected {
		t.Fatalf("got = %q, expected = %q", got, expected)
	}
}

func TestDFS_example2(t *testing.T) {
	nodesCount := 2
	edges := []edge{
		{1, 2},
	}
	startNode := 1

	expected := `1 2`

	var b bytes.Buffer
	dfs(&b, nodesCount, startNode, edges)
	got := b.String()

	if got != expected {
		t.Fatalf("got = %q, expected = %q", got, expected)
	}
}

func TestDFS_example3(t *testing.T) {
	nodesCount := 6
	edges := []edge{
		{4, 5},
		{5, 1},
		{1, 6},
		{5, 2},
		{4, 2},
		{1, 4},
		{4, 3},
	}
	startNode := 2

	expected := `2 4 1 5 6 3`

	var b bytes.Buffer
	dfs(&b, nodesCount, startNode, edges)
	got := b.String()

	if got != expected {
		t.Fatalf("got = %q, expected = %q", got, expected)
	}
}
