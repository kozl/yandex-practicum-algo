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

func dfs(w io.Writer, nodesCount, startNode int, edges []edge) {
	adjList := toAdjacencyList(nodesCount, edges)
	color := make([]int, nodesCount+1)
	visited := []int{}
	st := stack{}
	st.push(startNode)

	for !st.empty() {
		node := st.pop()
		if color[node] == white {
			visited = append(visited, node)
			color[node] = grey
			st.push(node)
			for _, adjacent := range adjList[node] {
				if color[adjacent] == white {
					st.push(adjacent)
				}
			}
		} else if color[node] == grey {
			color[node] = black
		}
	}
	fmt.Fprint(w, strings.Join(toString(visited), " "))
}

func toAdjacencyList(nodesCount int, edges []edge) [][]int {
	adjList := make([][]int, nodesCount+1)
	for _, edge := range edges {
		adjList[edge.from] = append(adjList[edge.from], edge.to)
		adjList[edge.to] = append(adjList[edge.to], edge.from)
	}
	for i := 0; i < len(adjList); i++ {
		sort.Sort(sort.Reverse(sort.IntSlice(adjList[i])))
	}
	return adjList
}

func toString(arr []int) []string {
	r := []string{}
	for _, i := range arr {
		r = append(r, fmt.Sprint(i))
	}
	return r
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
	startNode := readInt(scanner)
	dfs(os.Stdout, nodesCount, startNode, edges)
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

func readInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	stringInt := scanner.Text()
	res, _ := strconv.Atoi(stringInt)
	return res
}
