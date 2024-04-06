package main

import "testing"

func Test_solve(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "exampe 1",
			args: args{
				strs: []string{
					"abacaba",
					"abudabi",
					"abcdefg",
				},
			},
			want: 2,
		},
		{
			name: "example 2",
			args: args{
				strs: []string{
					"tutu",
					"kukuku",
				},
			},
			want: 0,
		},
		{
			name: "example 3",
			args: args{
				strs: []string{
					"qwe",
					"qwerty",
					"qwerpy",
				},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.strs); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
