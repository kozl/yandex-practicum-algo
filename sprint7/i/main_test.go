package main

import (
	"reflect"
	"testing"
)

func Test_solve(t *testing.T) {
	type args struct {
		n     int
		m     int
		field [][]int
	}
	tests := []struct {
		name          string
		args          args
		wantCollected int
		wantMoves     []move
	}{
		{
			name: "example 1",
			args: args{
				n: 2,
				m: 3,
				field: [][]int{
					{1, 1, 0},
					{1, 0, 1},
				},
			},
			wantCollected: 3,
			wantMoves:     []move{up, right, right},
		},
		{
			name: "example 2",
			args: args{
				n: 3,
				m: 3,
				field: [][]int{
					{0, 0, 1},
					{1, 1, 0},
					{1, 0, 0},
				},
			},
			wantCollected: 2,
			wantMoves:     []move{up, up, right, right},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCollected, gotMoves := solve(tt.args.n, tt.args.m, tt.args.field)
			if gotCollected != tt.wantCollected {
				t.Errorf("solve() gotCollected = %v, want %v", gotCollected, tt.wantCollected)
			}
			if !reflect.DeepEqual(gotMoves, tt.wantMoves) {
				t.Errorf("solve() gotMoves = %v, want %v", gotMoves, tt.wantMoves)
			}
		})
	}
}
