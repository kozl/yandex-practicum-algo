package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type edge struct {
	from int
	to   int
}

func isFull(nodes int, adjList []map[int]bool) bool {
	for node := 1; node <= nodes; node++ {
		if len(adjList[node]) != nodes-1 {
			return false
		}
	}
	return true
}

func solve(w io.Writer, nodesCount int, edges []edge) {
	adjList := toAdjacencyList(nodesCount, edges)
	yes := isFull(nodesCount, adjList)
	fmt.Fprintln(w, toStr(yes))
}

func toAdjacencyList(nodesCount int, edges []edge) []map[int]bool {
	adjList := make([]map[int]bool, nodesCount+1)
	for _, edge := range edges {
		if edge.from == edge.to {
			continue
		}
		if adjList[edge.from] == nil {
			adjList[edge.from] = map[int]bool{}
		}
		adjList[edge.from][edge.to] = true
		if adjList[edge.to] == nil {
			adjList[edge.to] = map[int]bool{}
		}
		adjList[edge.to][edge.from] = true

	}
	return adjList
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

func toStr(yes bool) string {
	if yes {
		return "YES"
	}
	return "NO"
}
