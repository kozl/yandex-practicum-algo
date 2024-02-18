package main

import (
	"bytes"
	"testing"
)

func Test_solve(t *testing.T) {
	type args struct {
		nodesCount int
		edges      []weightedEdge
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
				edges: []weightedEdge{
					{1, 2, 1},
					{2, 3, 3},
					{3, 4, 5},
					{1, 4, 2},
				},
			},
			wantW: `0 1 4 2
1 0 3 3
4 3 0 5
2 3 5 0
`,
		},
		{
			name: "example 2",
			args: args{
				nodesCount: 3,
				edges: []weightedEdge{
					{1, 2, 1},
					{1, 2, 2},
				},
			},
			wantW: `0 1 -1
1 0 -1
-1 -1 0
`,
		},
		{
			name: "example 3",
			args: args{
				nodesCount: 2,
				edges:      []weightedEdge{},
			},
			wantW: `0 -1
-1 0
`,
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
