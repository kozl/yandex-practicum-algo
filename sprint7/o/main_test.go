package main

import (
	"bytes"
	"testing"
)

func Test_dfs(t *testing.T) {
	type args struct {
		nodesCount int
		edges      []edge
		from       int
		to         int
	}
	tests := []struct {
		name  string
		args  args
		wantW string
	}{
		{
			name: "example 1",
			args: args{
				nodesCount: 3,
				edges: []edge{
					{1, 2},
					{1, 2},
					{2, 3},
				},
				from: 1,
				to:   3,
			},
			wantW: "2",
		},
		{
			name: "example 2",
			args: args{
				nodesCount: 5,
				edges: []edge{
					{1, 2},
					{3, 4},
					{4, 5},
				},
				from: 1,
				to:   5,
			},
			wantW: "0",
		},
		{
			name: "example 3",
			args: args{
				nodesCount: 3,
				edges: []edge{
					{1, 2},
					{2, 3},
					{1, 3},
				},
				from: 1,
				to:   1,
			},
			wantW: "1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			dfs(w, tt.args.nodesCount, tt.args.edges, tt.args.from, tt.args.to)
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("dfs() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}
