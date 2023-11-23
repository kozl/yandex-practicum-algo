package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSubsequence_Example1(t *testing.T) {
	s1 := "abc"
	s2 := "ahbgdcu"
	assert.True(t, isSubsequence(s1, s2))
}

func TestIsSubsequence_Example2(t *testing.T) {
	s1 := "abcp"
	s2 := "ahpc"
	assert.False(t, isSubsequence(s1, s2))
}

func TestIsSubsequence_Example3(t *testing.T) {
	s1 := "ijha"
	s2 := "hmrqvftefyixinahlzgbkidroxiptbbkjmtwpsujevkulgrjiwiwzyhngulrodiwyg"
	assert.False(t, isSubsequence(s1, s2))
}