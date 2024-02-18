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
				nodesCount: 3,
				edges: []edge{
					{1, 2},
					{2, 3},
				},
			},
			wantW: "YES\n",
		},
		{
			name: "example 2",
			args: args{
				nodesCount: 3,
				edges: []edge{
					{1, 2},
					{2, 3},
					{1, 3},
				},
			},
			wantW: "NO\n",
		},
		{
			name: "example 3",
			args: args{
				nodesCount: 100,
				edges: []edge{
					{2, 11},
					{4, 28},
					{7, 19},
					{8, 65},
					{10, 82},
					{10, 97},
					{12, 45},
					{12, 48},
					{14, 45},
					{14, 90},
					{16, 79},
					{16, 87},
					{20, 22},
					{21, 69},
					{23, 64},
					{23, 67},
					{24, 25},
					{24, 49},
					{25, 69},
					{28, 31},
					{28, 66},
					{29, 61},
					{30, 42},
					{31, 73},
					{32, 63},
					{32, 99},
					{36, 87},
					{37, 75},
					{38, 54},
					{38, 67},
					{38, 87},
					{41, 61},
					{42, 71},
					{43, 89},
					{45, 46},
					{45, 51},
					{45, 81},
					{46, 87},
					{47, 67},
					{48, 84},
					{50, 58},
					{50, 67},
					{50, 89},
					{51, 93},
					{52, 73},
					{54, 69},
					{54, 76},
					{61, 92},
					{71, 93},
					{76, 96},
					{78, 81},
					{81, 84},
					{84, 91},
					{90, 100},
				},
			},
			wantW: "NO\n",
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
