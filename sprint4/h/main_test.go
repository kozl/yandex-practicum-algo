package main

import "testing"

func Test_solve(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "example 1",
			args: args{
				a: "mxyskaoghi",
				b: "qodfrgmslc",
			},
			want: true,
		},
		{
			name: "example 2",
			args: args{
				a: "agg",
				b: "xdd",
			},
			want: true,
		},
		{
			name: "example 3",
			args: args{
				a: "agg",
				b: "xda",
			},
			want: false,
		},
		{
			name: "example 4",
			args: args{
				a: "aba",
				b: "xxx",
			},
			want: false,
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
