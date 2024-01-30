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

func printAdjacencyList(w io.Writer, nodesCount int, edges []edge) {
	var nodes []node
	for i := 0; i < nodesCount; i++ {
		nodes = append(nodes, node{
			id: i + 1,
		})
	}
	for _, edge := range edges {
		nodes[edge.from-1].adjacent = append(nodes[edge.from-1].adjacent, edge.to)
	}
	sort.Slice(nodes, func(i, j int) bool { return nodes[i].id < nodes[j].id })
	var builder strings.Builder
	for _, node := range nodes {
		builder.WriteString(fmt.Sprintf("%d", len(node.adjacent)))
		sort.Ints(node.adjacent)
		for _, nodeAdj := range node.adjacent {
			builder.WriteString(fmt.Sprintf(" %d", nodeAdj))
		}
		builder.WriteString("\n")
	}
	fmt.Fprint(w, builder.String())
}

type edge struct {
	from int
	to   int
}

type node struct {
	id       int
	adjacent []int
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
	printAdjacencyList(os.Stdout, nodesCount, edges)
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
