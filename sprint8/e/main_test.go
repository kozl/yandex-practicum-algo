package main

import "testing"

func Test_solve(t *testing.T) {
	type args struct {
		s          string
		operations []operation
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "example 1",
			args: args{
				s: "abacaba",
				operations: []operation{
					{s: "queue", pos: 2},
					{s: "deque", pos: 0},
					{s: "stack", pos: 7},
				},
			},
			want: "dequeabqueueacabastack",
		},
		{
			name: "example 2",
			args: args{
				s: "kukareku",
				operations: []operation{
					{s: "p", pos: 1},
					{s: "q", pos: 2},
				},
			},
			want: "kpuqkareku",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.s, tt.args.operations); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
