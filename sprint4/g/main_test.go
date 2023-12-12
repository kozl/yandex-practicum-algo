package main

import "testing"

func Test_solve(t *testing.T) {
	type args struct {
		array []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example 0",
			args: args{array: []int{0, 0, 1, 0, 1, 1, 1, 0, 0, 0}},
			want: 8,
		},
		{
			name: "example 1",
			args: args{array: []int{0, 1}},
			want: 2,
		},
		{
			name: "example 2",
			args: args{array: []int{0, 1, 0}},
			want: 2,
		},
		{
			name: "example 3",
			args: args{array: []int{0, 0, 1, 1, 1, 0, 0, 1, 1}},
			want: 8,
		},
		{
			name: "example 4",
			args: args{array: []int{1, 1, 1, 1, 0, 1, 1, 0, 0, 0, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 0, 0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0, 1, 0, 1, 1, 1, 0, 1, 1, 0, 0, 1, 0, 0, 0, 1, 0, 1, 0, 0, 0, 1, 1, 1, 1, 1, 1, 0, 1, 0, 0, 0}},
			want: 46,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.array); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
