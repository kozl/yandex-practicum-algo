package main

import (
	"testing"
)

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
				a: []int{1, 2, 3, 4, 5},
				b: []int{4, 5, 9},
			},
			want: 2,
		},
		{
			name: "example 3",
			args: args{
				a: []int{1},
				b: []int{1},
			},
			want: 1,
		},
		{
			name: "example 4",
			args: args{
				a: []int{69, 114, 84, 135, 195, 18, 64, 137, 164, 244, 81, 210, 107, 218, 49, 115, 208, 216, 206, 80, 241, 189, 173, 217, 50, 82, 202, 202, 222, 34},
				b: []int{214, 157, 57, 196, 58, 128, 89, 168, 197, 1, 200, 167, 243, 165, 89, 44, 18, 247, 175, 241, 136, 219, 114, 206, 254, 198, 65, 5, 144, 101},
			},
			want: 1,
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
