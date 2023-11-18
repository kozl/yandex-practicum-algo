package main

import "testing"

func TestFibMod(t *testing.T) {
	expect := 89
	got := fibMod(10, 2)
	if got != expect {
		t.Errorf("got %d, expected %d", got, expect)
	}
}
