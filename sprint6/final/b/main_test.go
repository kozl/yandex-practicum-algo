package main

import (
	"bytes"
	"testing"
)

func Test_solve(t *testing.T) {
	type args struct {
		cities         int
		mapDescription []string
	}
	tests := []struct {
		name  string
		args  args
		wantW string
	}{
		{
			name: "example 1",
			args: args{
				cities: 3,
				mapDescription: []string{
					"RB",
					"R",
				},
			},
			wantW: "NO\n",
		},
		{
			name: "example 2",
			args: args{
				cities: 4,
				mapDescription: []string{
					"BBB",
					"RB",
					"B",
				},
			},
			wantW: "YES\n",
		},
		{
			name: "example 3",
			args: args{
				cities: 5,
				mapDescription: []string{
					"RRRB",
					"BRR",
					"BR",
					"R",
				},
			},
			wantW: "NO\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			solve(w, tt.args.cities, tt.args.mapDescription)
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("solve() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}
