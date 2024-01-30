package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func printAdjacencyMatrix(w io.Writer, nodesCount int, edges []edge) {
	adjMatrix := initAdjMatrix(nodesCount)
	for _, edge := range edges {
		adjMatrix[edge.from-1][edge.to-1] = 1
	}

	for _, line := range adjMatrix {
		fmt.Fprintln(w, strings.Join(toStr(line), " "))
	}
}

func initAdjMatrix(nodesCount int) [][]int {
	matrix := make([][]int, nodesCount)
	for i := 0; i < nodesCount; i++ {
		matrix[i] = make([]int, nodesCount)
	}
	return matrix
}

func toStr(line []int) []string {
	result := make([]string, len(line))
	for idx, num := range line {
		result[idx] = fmt.Sprint(num)
	}
	return result
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
	printAdjacencyMatrix(os.Stdout, nodesCount, edges)
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
