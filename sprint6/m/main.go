package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	white = iota
	red
	black
)

type queue []int

func (s *queue) push(item int) {
	*s = append(*s, item)
}

func (s *queue) pop() int {
	popped := (*s)[0]
	*s = (*s)[1:]
	return popped
}

func (s *queue) empty() bool {
	return len(*s) == 0
}

type edge struct {
	from int
	to   int
}

func isBipartite(nodesCount int, adjacencyList [][]int) bool {
	color := make([]int, nodesCount+1)
	for node := 1; node <= nodesCount; node++ {
		if color[node] != white {
			continue
		}
		q := queue{}
		color[node] = red
		q.push(node)
		for !q.empty() {
			n := q.pop()
			for _, adjacent := range adjacencyList[n] {
				if color[adjacent] == white {
					color[adjacent] = oppositeColor(color[n])
					q.push(adjacent)
				} else if color[adjacent] != oppositeColor(color[n]) {
					return false
				}
			}
		}
	}
	return true
}

func oppositeColor(color int) int {
	if color == red {
		return black
	} else if color == black {
		return red
	}
	panic("unknown color")
}

func toAdjacencyList(nodesCount int, edges []edge) [][]int {
	adjList := make([][]int, nodesCount+1)
	for _, edge := range edges {
		adjList[edge.from] = append(adjList[edge.from], edge.to)
		adjList[edge.to] = append(adjList[edge.to], edge.from)
	}
	for i := 0; i < len(adjList); i++ {
		sort.Ints(adjList[i])
	}
	return adjList
}

func solve(w io.Writer, nodesCount int, edges []edge) {
	adjList := toAdjacencyList(nodesCount, edges)
	yes := isBipartite(nodesCount, adjList)
	fmt.Fprintln(w, toStr(yes))
}

func main() {
	scanner := makeScanner()
	raw := readArray(scanner)
	nodesCount, edgesCount := raw[0], raw[1]
	var edges []edge
	for i := 0; i < edgesCount; i++ {
		raw := readArray(scanner)
		edges = append(edges, edge{from: raw[0], to: raw[1]})
	}
	solve(os.Stdout, nodesCount, edges)
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 3 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func readArray(scanner *bufio.Scanner) []int {
	scanner.Scan()
	listString := strings.Split(scanner.Text(), " ")
	arr := make([]int, len(listString))
	for i := 0; i < len(listString); i++ {
		arr[i], _ = strconv.Atoi(listString[i])
	}
	return arr
}

func toStrSlice(ints []int) []string {
	s := make([]string, len(ints))
	for i := range ints {
		s[i] = fmt.Sprint(ints[i])
	}
	return s
}

func toStr(yes bool) string {
	if yes {
		return "YES"
	}
	return "NO"
}
