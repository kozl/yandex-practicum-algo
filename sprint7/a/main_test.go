package main

import "testing"

func Test_solve(t *testing.T) {
	type args struct {
		days   int
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example 1",
			args: args{
				days:   6,
				prices: []int{7, 1, 5, 3, 6, 4},
			},
			want: 7,
		},
		{
			name: "example 2",
			args: args{
				days:   5,
				prices: []int{1, 2, 3, 4, 5},
			},
			want: 4,
		},
		{
			name: "example 3",
			args: args{
				days:   6,
				prices: []int{1, 12, 12, 16, 1, 8},
			},
			want: 22,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.days, tt.args.prices); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
