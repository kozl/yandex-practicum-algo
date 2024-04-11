package main

import "testing"

func Test_solve(t *testing.T) {
	type args struct {
		line  string
		words []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "example 1",
			args: args{
				line: "examiwillpasstheexam",
				words: []string{
					"will",
					"pass",
					"the",
					"exam",
					"i",
				},
			},
			want: true,
		},
		{
			name: "example 2",
			args: args{
				line: "abacaba",
				words: []string{
					"abac",
					"caba",
				},
			},
			want: false,
		},
		{
			name: "example 3",
			args: args{
				line: "abacaba",
				words: []string{
					"abac",
					"caba",
					"aba",
				},
			},
			want: true,
		},
		{
			name: "example 4",
			args: args{
				line: "sscevscescescscsscevscevscesscsc",
				words: []string{
					"sce",
					"s",
					"scev",
					"sc",
				},
			},
			want: true,
		},
		{
			name: "example 5",
			args: args{
				line: "aaaaaaac",
				words: []string{
					"a",
					"aa",
					"aaa",
					"aaaa",
					"aaaaac",
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.line, tt.args.words); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
