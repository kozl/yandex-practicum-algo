package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	testcases := []struct {
		name  string
		docs  []string
		query string
		want  []int
	}{
		{
			name: "ex1_1",
			docs: []string{
				"i love coffee",
				"coffee with milk and sugar",
				"free tea for everyone",
			},
			query: "i like black coffee without milk",
			want:  []int{1, 2},
		},
		{
			name: "ex1_2",
			docs: []string{
				"i love coffee",
				"coffee with milk and sugar",
				"free tea for everyone",
			},
			query: "everyone loves new year",
			want:  []int{3},
		},
		{
			name: "ex1_3",
			docs: []string{
				"i love coffee",
				"coffee with milk and sugar",
				"free tea for everyone",
			},
			query: "mary likes black coffee without milk",
			want:  []int{2, 1},
		},
		{
			name: "ex2_1",
			docs: []string{
				"buy flat in moscow",
				"rent flat in moscow",
				"sell flat in moscow",
				"want flat in moscow like crazy",
				"clean flat in moscow on weekends",
				"renovate flat in moscow",
			},
			query: "flat in moscow for crazy weekends",
			want:  []int{4, 5, 1, 2, 3},
		},
		{
			name: "ex3_1",
			docs: []string{
				"i like dfs and bfs",
				"i like dfs dfs",
				"i like bfs with bfs and bfs",
			},
			query: "dfs dfs dfs dfs bfs",
			want:  []int{3, 1, 2},
		},
	}

	for _, tc := range testcases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			index := buildIndex(tc.docs)
			got := search(index, tc.query)
			assert.Equal(t, tc.want, got)
		})
	}
}
