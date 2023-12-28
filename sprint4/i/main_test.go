package main

import "testing"

func Test_solve(t *testing.T) {
	type args struct {
		a []int
		b []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example 1",
			args: args{
				a: []int{1, 2, 3, 2, 1},
				b: []int{3, 2, 1, 5, 6},
			},
			want: 3,
		},
		{
			name: "example 2",
			args: args{
				a: []int{1, 2, 3, 24, 5},
				b: []int{4, 5, 9},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
