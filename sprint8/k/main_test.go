package main

import "testing"

func Test_solve(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example 1",
			args: args{
				s: "gggggbbb",
				t: "bbef",
			},
			want: -1,
		},
		{
			name: "example 2",
			args: args{
				s: "z",
				t: "aaaaaaa",
			},
			want: 1,
		},
		{
			name: "example 3",
			args: args{
				s: "ccccz",
				t: "aaaaaz",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
