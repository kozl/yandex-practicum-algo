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
				nodesCount: 4,
				edges: []edge{
					{1, 2},
					{2, 3},
					{3, 4},
					{1, 4},
				},
				start: 3,
			},
			wantW: `3 2 4 1
`,
		},
		{
			name: "example 2",
			args: args{
				nodesCount: 2,
				edges: []edge{
					{2, 1},
				},
				start: 1,
			},
			wantW: `1 2
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			solve(w, tt.args.nodesCount, tt.args.start, tt.args.edges)
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("solve() = %q, want %q", gotW, tt.wantW)
			}
		})
	}
}
