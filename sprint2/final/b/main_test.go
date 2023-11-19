package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEval_Example1(t *testing.T) {
	s := newStack()
	w := bytes.Buffer{}
	eval(s, "2 1 + 3 *", &w)
	assert.Equal(t, "9\n", w.String())
}

func TestEval_Example2(t *testing.T) {
	s := newStack()
	w := bytes.Buffer{}
	eval(s, "7 2 + 4 * 2 +", &w)
	assert.Equal(t, "38\n", w.String())
}

func TestEval_FailedTest1(t *testing.T) {
	s := newStack()
	w := bytes.Buffer{}
	eval(s, "4 13 5 / +", &w)
	assert.Equal(t, "6\n", w.String())
}

func TestEval_FailedTest2(t *testing.T) {
	s := newStack()
	w := bytes.Buffer{}
	eval(s, "4 2 * 4 / 25 * 2 - 12 / 500 2 * + 2 / -999 + 71 + -1 * 2 / 1000 + 6 * 8065 - 787 + 66 * 456 - 45 * 10 /", &w)
	assert.Equal(t, "-2052\n", w.String())
}