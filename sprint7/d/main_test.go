package main

import "testing"

func Test_solve(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example 1",
			args: args{
				n: 5,
			},
			want: 8,
		},
		{
			name: "example 2",
			args: args{
				n: 2,
			},
			want: 2,
		},
		{
			name: "example 3",
			args: args{
				n: 10,
			},
			want: 89,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.n); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
