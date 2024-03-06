package main

import "testing"

func Test_solve(t *testing.T) {
	type args struct {
		total         int
		denominaitons []int
	}
	tests := []struct {
		name               string
		args               args
		wantBanknotesCount int
	}{
		{
			name: "example 1",
			args: args{
				total:         130,
				denominaitons: []int{10, 3, 40, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotBanknotesCount := solve(tt.args.total, tt.args.denominaitons); gotBanknotesCount != tt.wantBanknotesCount {
				t.Errorf("solve() = %v, want %v", gotBanknotesCount, tt.wantBanknotesCount)
			}
		})
	}
}
