package main

import (
	"bytes"
	"testing"
)

func Test_solve(t *testing.T) {
	type args struct {
		nodesCount int
		start      int
		edges      []edge
	}
	tests := []struct {
		name  string
		args  args
		wantW string
	}{
		{
			name: "example 1",
			args: args{
				nodesCount: 5,
				edges: []edge{
					{2, 1},
					{4, 5},
					{4, 3},
					{3, 2},
				},
				start: 2,
			},
			wantW: "3\n",
		},
		{
			name: "example 2",
			args: args{
				nodesCount: 3,
				edges: []edge{
					{3, 1},
					{1, 2},
					{2, 3},
				},
				start: 1,
			},
			wantW: "1\n",
		},
		{
			name: "example 3",
			args: args{
				nodesCount: 6,
				edges: []edge{
					{6, 1},
					{1, 3},
					{5, 1},
					{3, 5},
					{3, 4},
					{6, 5},
					{5, 2},
					{6, 2},
				},
				start: 4,
			},
			wantW: "3\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			solve(w, tt.args.nodesCount, tt.args.start, tt.args.edges)
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("solve() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}
