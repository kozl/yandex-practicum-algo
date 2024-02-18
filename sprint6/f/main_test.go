package main

import (
	"bytes"
	"testing"
)

func Test_solve(t *testing.T) {
	type args struct {
		nodesCount int
		start      int
		end        int
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
				start:      1,
				end:        5,
				edges: []edge{
					{2, 4},
					{3, 5},
					{2, 1},
					{2, 3},
					{4, 5},
				},
			},
			wantW: "3\n",
		},
		{
			name: "example 2",
			args: args{
				nodesCount: 4,
				start:      1,
				end:        3,
				edges: []edge{
					{2, 3},
					{4, 3},
					{2, 4},
				},
			},
			wantW: "-1\n",
		},
		{
			name: "example 3",
			args: args{
				nodesCount: 2,
				start:      1,
				end:        1,
				edges: []edge{
					{2, 1},
				},
			},
			wantW: "0\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			solve(w, tt.args.nodesCount, tt.args.start, tt.args.end, tt.args.edges)
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("solve() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}
