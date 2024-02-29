package main

import "testing"

func Test_solve(t *testing.T) {
	type args struct {
		n     int
		m     int
		field [][]int
	}
	tests := []struct {
		name string
		args args
		want int
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
			want: 3,
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
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.n, tt.args.m, tt.args.field); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
