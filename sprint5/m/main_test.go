package main

import "testing"

func TestSiftUp(t *testing.T) {
	sample := []int{-1, 12, 6, 8, 3, 15, 7}
	if siftUp(sample, 5) != 1 {
		t.Fatal("WA")
	}
}
