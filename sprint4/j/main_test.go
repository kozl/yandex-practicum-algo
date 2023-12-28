package main

import (
	"reflect"
	"testing"
)

func Test_solve(t *testing.T) {
	type args struct {
		arr []int
		s   int
	}
	tests := []struct {
		name string
		args args
		want [][4]int
	}{
		{
			name: "example 1",
			args: args{
				s:   10,
				arr: []int{2, 3, 2, 4, 1, 10, 3, 0},
			},
			want: [][4]int{
				{0, 3, 3, 4},
				{1, 2, 3, 4},
				{2, 2, 3, 3},
			},
		},
		{
			name: "example 2",
			args: args{
				s:   0,
				arr: []int{1, 0, -1, 0, 2, -2},
			},
			want: [][4]int{
				{-2, -1, 1, 2},
				{-2, 0, 0, 2},
				{-1, 0, 0, 1},
			},
		},
		{
			name: "example 3",
			args: args{
				s:   4,
				arr: []int{1, 1, 1, 1, 1},
			},
			want: [][4]int{
				{1, 1, 1, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.arr, tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
