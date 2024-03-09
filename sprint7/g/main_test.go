package main

import "testing"

func Test_solve(t *testing.T) {
	type args struct {
		n     int
		coins []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example 1",
			args: args{
				n:     5,
				coins: []int{3, 2, 1},
			},
			want: 5,
		},
		{
			name: "example 2",
			args: args{
				n:     3,
				coins: []int{2, 1},
			},
			want: 2,
		},
		{
			name: "example 3",
			args: args{
				n:     8,
				coins: []int{5},
			},
			want: 0,
		},
		{
			name: "example 4",
			args: args{
				n:     9,
				coins: []int{10, 2, 3, 7, 6, 5, 9},
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.n, tt.args.coins); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
