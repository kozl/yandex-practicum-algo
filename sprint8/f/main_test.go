package main

import "testing"

func Test_solve(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "example 1",
			args: args{
				words: []string{
					"caba",
					"aba",
					"caba",
					"abac",
					"aba",
				},
			},
			want: "aba",
		},
		{
			name: "example 2",
			args: args{
				words: []string{
					"b",
					"bc",
					"bcd",
				},
			},
			want: "b",
		},
		{
			name: "example 3",
			args: args{
				words: []string{
					"ciwlaxtnhhrnenw",
					"ciwnvsuni",
					"ciwaxeujmsmvpojqjkxk",
					"ciwnvsuni",
					"ciwnvsuni",
					"ciwuxlkecnofovq",
					"ciwuxlkecnofovq",
					"ciwodramivid",
					"ciwlaxtnhhrnenw",
					"ciwnvsuni",
				},
			},
			want: "ciwnvsuni",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.words); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
