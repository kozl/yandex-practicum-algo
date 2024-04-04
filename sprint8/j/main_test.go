package main

import (
	"reflect"
	"testing"
)

func Test_solve(t *testing.T) {
	type args struct {
		lines    []string
		patterns []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "example 1",
			args: args{
				lines: []string{
					"MamaMilaRamu",
					"MamaMia",
					"MonAmi",
				},
				patterns: []string{
					"MM",
					"MA",
				},
			},
			want: []string{"MamaMia", "MamaMilaRamu", "MonAmi"},
		},
		{
			name: "exapmle 2",
			args: args{
				lines: []string{
					"AlphaBetaGgamma",
					"AbcdBcdGggg",
				},
				patterns: []string{
					"ABGG",
					"ABG",
				},
			},
			want: []string{"AbcdBcdGggg", "AlphaBetaGgamma"},
		},
		{
			name: "example 3",
			args: args{
				lines: []string{
					"WudHnagkbhfwrbci",
					"WCUkvoxboxufsdap",
					"jdrxomezzrpuhbgi",
					"ZcGHdrPplfoldemu",
					"cylbtqwuxhiveznc",
				},
				patterns: []string{
					"WGHV",
					"NKVDT",
					"ZGHU",
				},
			},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.lines, tt.args.patterns); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
