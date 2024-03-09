package main

import (
	"reflect"
	"testing"
)

func Test_solve(t *testing.T) {
	type args struct {
		m     int
		items []item
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "example 1",
			args: args{
				m: 6,
				items: []item{
					{weight: 2, cost: 7},
					{weight: 4, cost: 2},
					{weight: 1, cost: 5},
					{weight: 2, cost: 1},
				},
			},
			want: []int{4, 3, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.m, tt.args.items); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
