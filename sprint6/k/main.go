package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

type weightedEdge struct {
	from   int
	to     int
	weight int
}

func dijkstra(nodesCount, start int, wAdjacencyList [][]weightedNode) []int {
	distances := make([]int, nodesCount+1)
	visited := make([]bool, nodesCount+1)
	for i := 1; i <= nodesCount; i++ {
		distances[i] = math.MaxInt32
	}
	distances[start] = 0

	for node := start; node != -1; node = getNextNode(distances, visited) {
		visited[node] = true
		for _, adjacent := range wAdjacencyList[node] {
			relax(node, adjacent, distances)
		}
	}
	return distances[1:]
}

func relax(node int, adjacent weightedNode, distances []int) {
	if distances[node]+adjacent.weight < distances[adjacent.node] {
		distances[adjacent.node] = distances[node] + adjacent.weight
	}
}

func getNextNode(distances []int, visited []bool) int {
	currMinimum := math.MaxInt32
	currNode := -1
	for node, dist := range distances[1:] {
		node++
		if !visited[node] && currMinimum > dist {
			currNode = node
			currMinimum = dist
		}
	}
	return currNode
}

type weightedNode struct {
	node   int
	weight int
}

func toWadjacencyList(nodesCount int, edges []weightedEdge) [][]weightedNode {
	adjList := make([][]weightedNode, nodesCount+1)
	for _, edge := range edges {
		adjList[edge.from] = append(adjList[edge.from], weightedNode{node: edge.to, weight: edge.weight})
		adjList[edge.to] = append(adjList[edge.to], weightedNode{node: edge.from, weight: edge.weight})
	}
	return adjList
}

func solve(w io.Writer, nodesCount int, edges []weightedEdge) {
	adjList := toWadjacencyList(nodesCount, edges)
	for node := 1; node <= nodesCount; node++ {
		distances := dijkstra(nodesCount, node, adjList)
		fmt.Fprintln(w, strings.Join(toStrSlice(distances), " "))
	}
}

func main() {
	scanner := makeScanner()
	raw := readArray(scanner)
	nodesCount, edgesCount := raw[0], raw[1]
	var edges []weightedEdge
	for i := 0; i < edgesCount; i++ {
		raw := readArray(scanner)
		edges = append(edges, weightedEdge{from: raw[0], to: raw[1], weight: raw[2]})
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
		if ints[i] != math.MaxInt32 {
			s[i] = fmt.Sprint(ints[i])
		} else {
			s[i] = "-1"
		}
	}
	return s
}
