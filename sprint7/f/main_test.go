package main

import "testing"

func Test_solve(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example 1",
			args: args{
				n: 6,
				k: 3,
			},
			want: 13,
		},
		{
			name: "example 2",
			args: args{
				n: 7,
				k: 7,
			},
			want: 32,
		},
		{
			name: "example 3",
			args: args{
				n: 2,
				k: 2,
			},
			want: 1,
		},
		{
			name: "example 4",
			args: args{
				n: 62,
				k: 44,
			},
			want: 535806680,
		},
		{
			name: "example 5",
			args: args{
				n: 79,
				k: 34,
			},
			want: 470472762,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
