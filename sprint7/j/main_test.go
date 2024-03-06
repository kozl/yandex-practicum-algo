package main

import (
	"reflect"
	"testing"
)

func Test_solve(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name        string
		args        args
		wantIndices []int
	}{
		{
			name: "example 1",
			args: args{
				a: []int{4, 2, 9, 1, 13},
			},
			wantIndices: []int{1, 3, 5},
		},
		{
			name: "example 2",
			args: args{
				a: []int{1, 2, 4, 8, 16, 32},
			},
			wantIndices: []int{1, 2, 3, 4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotIndices := solve(tt.args.a)
			if !reflect.DeepEqual(gotIndices, tt.wantIndices) {
				t.Errorf("solve() gotIndices = %v, want %v", gotIndices, tt.wantIndices)
			}
		})
	}
}
