package main

import (
	"bytes"
	"io"
	"reflect"
	"strings"
	"testing"
)

func Test_solve(t *testing.T) {
	type args struct {
		lessons []lesson
	}
	tests := []struct {
		name string
		args args
		want []lesson
	}{
		{
			name: "example 1",
			args: args{
				lessons: []lesson{
					newLesson("9", "10"),
					newLesson("9.3", "10.3"),
					newLesson("10", "11"),
					newLesson("10.3", "11.3"),
					newLesson("11", "12"),
				},
			},
			want: []lesson{
				newLesson("9", "10"),
				newLesson("10", "11"),
				newLesson("11", "12"),
			},
		},
		{
			name: "example 2",
			args: args{
				lessons: []lesson{
					newLesson("9", "10"),
					newLesson("11", "12.25"),
					newLesson("12.15", "13.3"),
				},
			},
			want: []lesson{
				newLesson("9", "10"),
				newLesson("11", "12.25"),
			},
		},
		{
			name: "example 3",
			args: args{
				lessons: []lesson{
					newLesson("19", "19"),
					newLesson("7", "14"),
					newLesson("12", "14"),
					newLesson("8", "22"),
					newLesson("22", "23"),
					newLesson("5", "21"),
					newLesson("9", "23"),
				},
			},
			want: []lesson{
				newLesson("7", "14"),
				newLesson("19", "19"),
				newLesson("22", "23"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.lessons); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_run(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name  string
		args  args
		wantW string
	}{
		{
			name: "example 4",
			args: args{
				strings.NewReader(`59
15 22
17 20
12 13
21 23
15 15
3 23
20 23
7 18
11 13
2 16
7 19
1 10
16 23
15 17
15 19
12 14
8 9
8 17
19 23
12 15
3 10
3 8
17 20
20 21
0 0
17 21
13 17
2 23
20 20
18 19
9 10
7 10
23 23
22 22
8 10
4 9
21 21
18 22
14 19
19 20
22 23
12 22
3 9
15 23
2 21
8 8
10 15
13 13
0 7
11 19
0 22
2 6
15 16
5 8
20 23
18 23
11 22
17 20
12 14`),
			},
			wantW: `17
0 0
2 6
8 8
8 9
9 10
11 13
13 13
15 15
15 16
18 19
19 20
20 20
20 21
21 21
22 22
22 23
23 23
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			run(tt.args.r, w)
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("run() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}
