package main

import "testing"

func Test_solve(t *testing.T) {
	type args struct {
		n    int
		m    int
		gold []int
	}
	tests := []struct {
		name          string
		args          args
		wantMaxWeight int
	}{
		{
			name: "example 1",
			args: args{
				n:    5,
				m:    15,
				gold: []int{3, 8, 1, 2, 5},
			},
			wantMaxWeight: 15,
		},
		{
			name: "example 2",
			args: args{
				n:    5,
				m:    19,
				gold: []int{10, 10, 7, 7, 4},
			},
			wantMaxWeight: 18,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotMaxWeight := solve(tt.args.n, tt.args.m, tt.args.gold); gotMaxWeight != tt.wantMaxWeight {
				t.Errorf("solve() = %v, want %v", gotMaxWeight, tt.wantMaxWeight)
			}
		})
	}
}
