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
	grey
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

func bfsShortestDistance(nodesCount, start, end int, adjacencyList [][]int) int {
	color := make([]int, nodesCount+1)
	q := queue{}
	distance := make([]int, nodesCount+1)
	color[start] = grey
	distance[start] = 0
	q.push(start)
	for !q.empty() {
		n := q.pop()
		if n == end {
			return distance[n]
		}
		for _, adjacent := range adjacencyList[n] {
			if color[adjacent] == white {
				color[adjacent] = grey
				distance[adjacent] = distance[n] + 1
				q.push(adjacent)
			}
		}
		color[n] = black
	}
	return -1
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

func solve(w io.Writer, nodesCount, start, end int, edges []edge) {
	adjList := toAdjacencyList(nodesCount, edges)
	distance := bfsShortestDistance(nodesCount, start, end, adjList)
	fmt.Fprintln(w, distance)
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
	raw = readArray(scanner)
	start, end := raw[0], raw[1]
	solve(os.Stdout, nodesCount, start, end, edges)
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
