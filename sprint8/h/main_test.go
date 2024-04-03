package main

import "testing"

func Test_solve(t *testing.T) {
	type args struct {
		line        string
		pattern     string
		replacement string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "example 1",
			args: args{
				line:        "pingpong",
				pattern:     "ng",
				replacement: "mpi",
			},
			want: "pimpipompi",
		},
		{
			name: "example 2",
			args: args{
				line:        "aaa",
				pattern:     "a",
				replacement: "ab",
			},
			want: "ababab",
		},
		{
			name: "example 3",
			args: args{
				line:        "soxiogffiogffqxiogffiogffbqqtl",
				pattern:     "iogff",
				replacement: "bcfp",
			},
			want: "soxbcfpbcfpqxbcfpbcfpbqqtl",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.line, tt.args.pattern, tt.args.replacement); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
