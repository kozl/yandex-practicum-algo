package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPolyHash(t *testing.T) {
	testcases := []struct {
		a    int
		s    string
		mod  int
		want int
	}{
		{
			a:    123,
			mod:  100003,
			s:    "a",
			want: 97,
		},
		{
			a:    123,
			mod:  100003,
			s:    "hash",
			want: 6080,
		},
		{
			a:    123,
			mod:  100003,
			s:    "HaSH",
			want: 56156,
		},
	}

	for _, tc := range testcases {
		got := polyHash(tc.s, tc.a, tc.mod)
		assert.Equal(t, tc.want, got)
	}
}
