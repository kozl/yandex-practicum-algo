package main

import (
	"reflect"
	"testing"
)

func Test_solve(t *testing.T) {
	type args struct {
		a []int
		b []int
	}
	tests := []struct {
		name     string
		args     args
		wantS    int
		wantAIdx []int
		wantBIdx []int
	}{
		{
			name: "example 1",
			args: args{
				a: []int{4, 9, 2, 4, 6},
				b: []int{9, 4, 0, 0, 2, 8, 4},
			},
			wantS:    3,
			wantAIdx: []int{1, 3, 4},
			wantBIdx: []int{2, 5, 7},
		},
		{
			name: "example 2",
			args: args{
				a: []int{1, 1, 1, 1},
				b: []int{2, 2},
			},
			wantS:    0,
			wantAIdx: nil,
			wantBIdx: nil,
		},
		{
			name: "example 3",
			args: args{
				a: []int{1, 2, 1, 9, 1, 2, 1, 9},
				b: []int{9, 9, 1, 9, 9},
			},
			wantS:    3,
			wantAIdx: []int{3, 4, 8},
			wantBIdx: []int{3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotS, gotAIdx, gotBIdx := solve(tt.args.a, tt.args.b)
			if gotS != tt.wantS {
				t.Errorf("solve() gotS = %v, want %v", gotS, tt.wantS)
			}
			if !reflect.DeepEqual(gotAIdx, tt.wantAIdx) {
				t.Errorf("solve() gotAIdx = %v, want %v", gotAIdx, tt.wantAIdx)
			}
			if !reflect.DeepEqual(gotBIdx, tt.wantBIdx) {
				t.Errorf("solve() gotBIdx = %v, want %v", gotBIdx, tt.wantBIdx)
			}
		})
	}
}
