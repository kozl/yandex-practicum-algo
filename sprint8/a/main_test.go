package main

import "testing"

func Test_solve(t *testing.T) {
	type args struct {
		sentence string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "example 1",
			args: args{
				sentence: "one two three",
			},
			want: "three two one",
		},
		{
			name: "example 2",
			args: args{
				sentence: "hello",
			},
			want: "hello",
		},
		{
			name: "example 3",
			args: args{
				sentence: "may the force be with you",
			},
			want: "you with be force the may",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.sentence); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
