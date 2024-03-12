package main

import (
	"testing"
)

func Test_countPaths(t *testing.T) {
	type args struct {
		nodesCount int
		edges      []edge
		from       int
		to         int
	}
	tests := []struct {
		name string
		args args
		want int
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
			want: 2,
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
			want: 0,
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
			want: 1,
		},
		{
			name: "example 4",
			args: args{
				nodesCount: 10,
				edges: []edge{
					{5, 4},
					{2, 10},
					{2, 10},
					{2, 3},
					{4, 1},
					{5, 1},
					{10, 6},
					{4, 1},
					{1, 6},
					{2, 10},
					{4, 6},
					{9, 6},
					{9, 10},
					{2, 7},
					{4, 1},
					{9, 1},
					{7, 3},
					{2, 9},
					{10, 4},
					{2, 3},
				},
				from: 9,
				to:   1,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countPaths(tt.args.nodesCount, tt.args.edges, tt.args.from, tt.args.to)
			if got != tt.want {
				t.Errorf("countPaths() = %v, want %v", got, tt.want)
			}
		})
	}
}
