package main

import "testing"

func Test_levenstein(t *testing.T) {
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
				s: "abacaba",
				t: "abaabc",
			},
			want: 2,
		},
		{
			name: "example 2",
			args: args{
				s: "innokentiy",
				t: "innnokkentia",
			},
			want: 3,
		},
		{
			name: "example 3",
			args: args{
				s: "r",
				t: "x",
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := levenstein(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("levenstein() = %v, want %v", got, tt.want)
			}
		})
	}
}
