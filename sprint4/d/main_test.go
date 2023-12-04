package main

import (
	"reflect"
	"testing"
)

func Test_uniq(t *testing.T) {
	type args struct {
		array []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "1 пример",
			args: args{
				array: []string{
					"вышивание крестиком",
					"рисование мелками на парте",
					"настольный керлинг",
					"настольный керлинг",
					"кухня африканского племени ужасмай",
					"тяжелая атлетика",
					"таракановедение",
					"таракановедение",
				},
			},
			want: []string{
				"вышивание крестиком",
				"рисование мелками на парте",
				"настольный керлинг",
				"кухня африканского племени ужасмай",
				"тяжелая атлетика",
				"таракановедение",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniq(tt.args.array); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("uniq() = %v, want %v", got, tt.want)
			}
		})
	}
}
