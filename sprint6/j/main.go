package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

const (
	white = iota
	grey
	black
)

type stack []int

func (s *stack) push(item int) {
	*s = append(*s, item)
}

func (s *stack) pop() int {
	popped := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return popped
}

func (s *stack) empty() bool {
	return len(*s) == 0
}

type edge struct {
	from int
	to   int
}

func topologicSort(nodesCount int, adjacencyList [][]int) []int {
	color := make([]int, nodesCount+1)
	st := stack{}
	result := stack{}
	for node := 1; node < len(color); node++ {
		if color[node] != white {
			continue
		}
		st.push(node)
		for !st.empty() {
			node := st.pop()
			if color[node] == white {
				color[node] = grey
				st.push(node)
				for _, adjacent := range adjacencyList[node] {
					if color[adjacent] == white {
						st.push(adjacent)
					}
				}
			} else if color[node] == grey {
				color[node] = black
				result.push(node)
			}
		}
	}
	slices.Reverse(result)
	return result
}

func toAdjacencyList(nodesCount int, edges []edge) [][]int {
	adjList := make([][]int, nodesCount+1)
	for _, edge := range edges {
		adjList[edge.from] = append(adjList[edge.from], edge.to)
	}
	for i := 0; i < len(adjList); i++ {
		sort.Sort(sort.Reverse(sort.IntSlice(adjList[i])))
	}
	return adjList
}

func solve(w io.Writer, nodesCount int, edges []edge) {
	adjList := toAdjacencyList(nodesCount, edges)
	nodes := topologicSort(nodesCount, adjList)
	fmt.Fprintln(w, strings.Join(toStrSlice(nodes), " "))
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
