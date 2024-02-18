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

func bfs(nodesCount, start int, adjacencyList [][]int) []int {
	color := make([]int, nodesCount+1)
	q := queue{}
	visited := []int{}
	color[start] = grey
	q.push(start)
	for !q.empty() {
		n := q.pop()
		for _, adjacent := range adjacencyList[n] {
			if color[adjacent] == white {
				color[adjacent] = grey
				q.push(adjacent)
			}
		}
		color[n] = black
		visited = append(visited, n)
	}
	return visited
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

func solve(w io.Writer, nodesCount, start int, edges []edge) {
	adjList := toAdjacencyList(nodesCount, edges)
	visited := bfs(nodesCount, start, adjList)
	fmt.Fprintln(w, strings.Join(toStrSlice(visited), " "))
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
	start := readInt(scanner)
	solve(os.Stdout, nodesCount, start, edges)
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

func readInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	stringInt := scanner.Text()
	res, _ := strconv.Atoi(stringInt)
	return res
}
