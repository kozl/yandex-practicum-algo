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
		name string
		args args
		want string
	}{
		{
			name: "example 1",
			args: args{
				nodesCount: 5,
				edges: []edge{
					{from: 3, to: 2},
					{from: 3, to: 4},
					{from: 2, to: 5},
				},
			},
			want: "3 4 2 5 1\n",
		},
		{
			name: "example 2",
			args: args{
				nodesCount: 6,
				edges: []edge{
					{from: 6, to: 4},
					{from: 4, to: 1},
					{from: 5, to: 1},
				},
			},
			want: "6 5 4 3 2 1\n",
		},
		{
			name: "example 3",
			args: args{
				nodesCount: 4,
				edges:      []edge{},
			},
			want: "4 3 2 1\n",
		},
		{
			name: "example 4",
			args: args{
				nodesCount: 6,
				edges:      []edge{
					{1, 6},
					{5, 6},
					{1, 3},
					{1, 2},
					{2, 4},
					{4, 3},
				},
			},
			want: "1 2 4 3 5 6\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			solve(w, tt.args.nodesCount, tt.args.edges)
			if gotW := w.String(); gotW != tt.want {
				t.Errorf("solve() = %q, want %q", gotW, tt.want)
			}
		})
	}
}
