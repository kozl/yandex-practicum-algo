package main

import (
	"bytes"
	"testing"
)

func Test_solve(t *testing.T) {
	type args struct {
		nodesCount int
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
				nodesCount: 4,
				edges: []edge{
					{1, 2, 5},
					{1, 3, 6},
					{2, 4, 8},
					{3, 4, 3},
				},
			},
			wantW: "19\n",
		},
		{
			name: "example 2",
			args: args{
				nodesCount: 3,
				edges: []edge{
					{1, 2, 1},
					{1, 2, 2},
					{2, 3, 1},
				},
			},
			wantW: "3\n",
		},
		{
			name: "example 3",
			args: args{
				nodesCount: 2,
				edges:      []edge{},
			},
			wantW: "Oops! I did it again\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			solve(w, tt.args.nodesCount, tt.args.edges)
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("solve() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}
