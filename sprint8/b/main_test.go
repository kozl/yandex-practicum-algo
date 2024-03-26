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
		want bool
	}{
		{
			name: "example 1",
			args: args{
				s: "abcdefg",
				t: "abdefg",
			},
			want: true,
		},
		{
			name: "example 2",
			args: args{
				s: "helo",
				t: "hello",
			},
			want: true,
		},
		{
			name: "example 3",
			args: args{
				s: "dog",
				t: "fog",
			},
			want: true,
		},
		{
			name: "example 4",
			args: args{
				s: "mama",
				t: "papa",
			},
			want: false,
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
