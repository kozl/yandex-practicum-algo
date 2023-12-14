package main

import "testing"

func Test_polyHashSubstring(t *testing.T) {
	type args struct {
		start int
		end   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				start: 1,
				end:   1,
			},
			want: 97,
		},
		{
			name: "2",
			args: args{
				start: 1,
				end:   5,
			},
			want: 225076,
		},
		{
			name: "3",
			args: args{
				start: 2,
				end:   3,
			},
			want: 98099,
		},
		{
			name: "4",
			args: args{
				start: 3,
				end:   4,
			},
			want: 99100,
		},
		{
			name: "5",
			args: args{
				start: 4,
				end:   4,
			},
			want: 100,
		},
		{
			name: "6",
			args: args{
				start: 1,
				end:   8,
			},
			want: 436420,
		},
		{
			name: "7",
			args: args{
				start: 5,
				end:   8,
			},
			want: 193195,
		},
	}
	a := 1000
	mod := 1000009
	s := "abcdefgh"
	pows := precalculatePows(a, mod, len(s))
	prefixHashes := precalculatePrefixHashes(s, a, mod)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := polyHashSubstring(prefixHashes, pows, tt.args.start, tt.args.end, a, mod); got != tt.want {
				t.Errorf("polyHashSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
