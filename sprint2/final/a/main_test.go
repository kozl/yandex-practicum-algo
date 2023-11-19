package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDequeue_Basic1(t *testing.T) {
	d := newDeq(5)
	for i := 0; i < 5; i++ {
		err := d.pushBack(i)
		assert.NoError(t, err)
	}
	for i := 0; i < 5; i++ {
		v, err := d.popFront()
		assert.NoError(t, err)
		assert.Equal(t, i, v)
	}
	_, err := d.popFront()
	assert.Error(t, err)
}

func TestDequeue_Basic2(t *testing.T) {
	d := newDeq(5)
	for i := 0; i < 5; i++ {
		err := d.pushFront(i)
		assert.NoError(t, err)
	}
	for i := 0; i < 5; i++ {
		v, err := d.popBack()
		assert.NoError(t, err)
		assert.Equal(t, i, v)
	}
	_, err := d.popBack()
	assert.Error(t, err)
}

func TestDequeue_Basic3(t *testing.T) {
	d := newDeq(5)
	for i := 0; i < 5; i++ {
		err := d.pushFront(i)
		assert.NoError(t, err)
	}
	for i := 4; i >= 0; i-- {
		v, err := d.popFront()
		assert.NoError(t, err)
		assert.Equal(t, i, v)
	}
	_, err := d.popFront()
	assert.Error(t, err)
}

func TestEval_Example1(t *testing.T) {
	d := newDeq(4)
	var w bytes.Buffer

	eval(d, "push_front 861", &w)
	eval(d, "push_front -819", &w)
	eval(d, "pop_back", &w)
	eval(d, "pop_back", &w)

	assert.Equal(t, "861\n-819\n", w.String())
}

func TestEval_Example2(t *testing.T) {
	d := newDeq(10)
	var w bytes.Buffer

	eval(d, "push_front -855", &w)
	eval(d, "push_front 0", &w)
	eval(d, "pop_back", &w)
	eval(d, "pop_back", &w)
	eval(d, "push_back 844", &w)
	eval(d, "pop_back", &w)
	eval(d, "push_back 823", &w)

	assert.Equal(t, "-855\n0\n844\n", w.String())
}

func TestEval_Example3(t *testing.T) {
	d := newDeq(6)
	var w bytes.Buffer

	eval(d, "push_front -201", &w)
	eval(d, "push_back 959", &w)
	eval(d, "push_back 102", &w)
	eval(d, "push_front 20", &w)
	eval(d, "pop_front", &w)
	eval(d, "pop_back", &w)

	assert.Equal(t, "20\n102\n", w.String())
}
