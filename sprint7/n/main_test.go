package main

import (
	"reflect"
	"testing"
)

func Test_solve(t *testing.T) {
	type args struct {
		days []int
	}
	tests := []struct {
		name string
		args args
		want state
	}{
		{
			name: "example 1",
			args: args{
				days: []int{500, 501, 300},
			},
			want: state{
				sum:      1001,
				freeDays: []int{3},
			},
		},
		{
			name: "example 2",
			args: args{
				days: []int{502, 501, 503, 504},
			},
			want: state{
				sum:      1003,
				freeDays: []int{3, 4},
			},
		},
		{
			name: "example 3",
			args: args{
				days: []int{501},
			},
			want: state{
				sum:      501,
				freeDays: []int{},
			},
		},
		{
			name: "example 4",
			args: args{
				days: []int{149, 80, 275, 832, 462, 567, 168, 79, 890, 144, 154, 609, 554, 562, 416, 182, 643, 366, 227, 917, 208, 975, 365, 32, 971, 447, 662, 15, 652, 665, 33, 975, 949, 58, 21, 288, 883, 673, 637, 443, 317, 211, 33, 869, 647, 536, 297, 867, 440, 438, 280, 836, 116, 11, 259, 618, 738, 452, 630, 628, 996, 987, 392, 917, 35, 377, 857, 989, 631, 163, 531, 780, 853, 980, 684, 856, 288, 112, 662, 717, 712, 89, 996, 578, 749, 326, 107, 739, 18, 260, 484, 263, 508, 226, 928, 446, 921, 149, 954, 780},
			},
			want: state{
				sum:      27084,
				freeDays: []int{5, 9, 20, 22, 25, 26, 32, 33, 37, 40, 44, 48, 49, 58, 61, 62, 64, 68, 72, 73, 74, 76, 83, 91, 95, 96, 97, 99, 100},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.days); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
