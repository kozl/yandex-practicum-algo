package main

import "testing"

func Test_solve(t *testing.T) {
	type args struct {
		lines []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "example 1",
			args: args{
				lines: []string{
					"2[a]2[ab]",
					"3[a]2[r2[t]]",
					"a2[aa3[b]]",
				},
			},
			want: "aaa",
		},
		{
			name: "example 2",
			args: args{
				lines: []string{
					"abacabaca",
					"2[abac]a",
					"3[aba]",
				},
			},
			want: "aba",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.lines); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
