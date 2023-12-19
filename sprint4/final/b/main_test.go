package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap_Example1(t *testing.T) {
	m := newMap(2 << 20)
	w := &bytes.Buffer{}

	commands := []string{
		"get 1",
		"put 1 10",
		"put 2 4",
		"get 1",
		"get 2",
		"delete 2",
		"get 2",
		"put 1 5",
		"get 1",
		"delete 2",
	}

	expectedOutput := `None
10
4
4
None
5
None
`

	for _, cmd := range commands {
		eval(m, cmd, w)
	}

	assert.Equal(t, expectedOutput, w.String())
}

func TestMap_Example2(t *testing.T) {
	m := newMap(2 << 20)
	w := &bytes.Buffer{}

	commands := []string{
		"get 9",
		"delete 9",
		"put 9 1",
		"get 9",
		"put 9 2",
		"get 9",
		"put 9 3",
		"get 9",
	}

	expectedOutput := `None
None
1
2
3
`

	for _, cmd := range commands {
		eval(m, cmd, w)
	}

	assert.Equal(t, expectedOutput, w.String())
}

func TestMap_Example3(t *testing.T) {
	m := newMap(2 << 20)
	w := &bytes.Buffer{}

	commands := []string{
		"put 20 27",
		"get 20",
		"put 20 21",
		"get 20",
		"get 20",
		"get -1",
		"get 20",
		"get -3",
		"delete 20",
		"get -29",
		"get -33",
		"delete -29",
		"get 16",
		"get 14",
		"put 29 39",
	}

	expectedOutput := `27
21
21
None
21
None
21
None
None
None
None
None
`

	for _, cmd := range commands {
		eval(m, cmd, w)
	}

	assert.Equal(t, expectedOutput, w.String())
}
