package main

import "testing"

func TestSiftDown(t *testing.T) {
	sample := []int{-1, 12, 1, 8, 3, 4, 7}
	if siftDown(sample, 2) != 5 {
		t.Fatal("WA")
	}
}
