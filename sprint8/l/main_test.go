package main

import (
	"reflect"
	"testing"
)

func Test_solve(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "example 1",
			args: args{
				s: "abracadabra",
			},
			want: []int{0, 0, 0, 1, 0, 1, 0, 1, 2, 3, 4},
		},
		{
			name: "example 2",
			args: args{
				s: "xxzzxxz",
			},
			want: []int{0, 1, 0, 0, 1, 2, 3},
		},
		{
			name: "example 2",
			args: args{
				s: "aaaaa",
			},
			want: []int{0, 1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
