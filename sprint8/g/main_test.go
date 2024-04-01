package main

import (
	"reflect"
	"testing"
)

func Test_solve(t *testing.T) {
	type args struct {
		x []int
		a []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "example 1",
			args: args{
				x: []int{3, 9, 1, 2, 5, 10, 9, 1, 7},
				a: []int{4, 10},
			},
			want: []int{1, 8},
		},
		{
			name: "example 2",
			args: args{
				x: []int{1, 2, 3, 4, 5},
				a: []int{10, 11, 12},
			},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.x, tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
