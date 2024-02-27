package main

import "testing"

func Test_solve(t *testing.T) {
	type args struct {
		capacity int
		piles    []goldenPile
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example 1",
			args: args{
				capacity: 10,
				piles: []goldenPile{
					{pricePerKg: 8, weightKg: 1},
					{pricePerKg: 2, weightKg: 10},
					{pricePerKg: 4, weightKg: 5},
				},
			},
			want: 36,
		},
		{
			name: "example 2",
			args: args{
				capacity: 10_000,
				piles: []goldenPile{
					{pricePerKg: 4, weightKg: 20},
				},
			},
			want: 80,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.capacity, tt.args.piles); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
