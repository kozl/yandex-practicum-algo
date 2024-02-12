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

func dfs(nodesCount, startNode int, edges []edge) [][2]int {
	adjList := toAdjacencyList(nodesCount, edges)
	color := make([]int, nodesCount+1)
	times := make([][2]int, nodesCount+1)
	time := 0
	st := stack{}
	st.push(startNode)

	for !st.empty() {
		node := st.pop()
		if color[node] == white {
			times[node][0] = time
			color[node] = grey
			time++
			st.push(node)
			for _, adjacent := range adjList[node] {
				if color[adjacent] == white {
					st.push(adjacent)
				}
			}
		} else if color[node] == grey {
			times[node][1] = time
			color[node] = black
			time++
		}
	}
	return times
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
	times := dfs(nodesCount, 1, edges)
	for i := 1; i <= nodesCount; i++ {
		fmt.Fprintf(w, "%d %d\n", times[i][0], times[i][1])
	}
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

func readInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	stringInt := scanner.Text()
	res, _ := strconv.Atoi(stringInt)
	return res
}
