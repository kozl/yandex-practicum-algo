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
				nodesCount: 6,
				edges: []edge{
					{1, 2},
					{6, 5},
					{2, 3},
				},
			},
			wantW: `3
1 2 3
4
5 6
`,
		},
		{
			name: "example 2",
			args: args{
				nodesCount: 2,
				edges:      []edge{},
			},
			wantW: `2
1
2
`,
		},
		{
			name: "example 3",
			args: args{
				nodesCount: 4,
				edges: []edge{
					{2, 3},
					{2, 1},
					{4, 3},
				},
			},
			wantW: `1
1 2 3 4
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			solve(w, tt.args.nodesCount, tt.args.edges)
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("solve() = %q, want %q", gotW, tt.wantW)
			}
		})
	}
}
