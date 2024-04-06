package main

import "testing"

func Test_solve(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "example 1",
			args: args{
				s: "aaaabb",
			},
			want: "aabbaa",
		},
		{
			name: "example 2",
			args: args{
				s: "pabcd",
			},
			want: "a",
		},
		{
			name: "example 3",
			args: args{
				s: "aaabbb",
			},
			want: "ababa",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.s); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
