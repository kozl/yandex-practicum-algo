package main

import "testing"

func Test_solve(t *testing.T) {
	type args struct {
		scores []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "example 1",
			args: args{
				scores: []int{1, 5, 7, 1},
			},
			want: true,
		},
		{
			name: "example 2",
			args: args{
				scores: []int{2, 10, 9},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.scores); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
