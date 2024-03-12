package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	white = iota
	grey
	black
)

var C = int(math.Pow(float64(10), float64(9))) + 7

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

func countPaths(nodesCount int, edges []edge, from, to int) int {
	revAdjList := toReverseAdjacencyList(nodesCount, edges)
	dp := make([]int, nodesCount+1)
	color := make([]int, nodesCount+1)
	s := stack{}
	dp[from] = 1
	color[from] = black
	s.push(to)
	for !s.empty() {
		n := s.pop()
		if color[n] == white {
			s.push(n)
			color[n] = grey
			for _, revAdj := range revAdjList[n] {
				if color[revAdj] == white {
					s.push(revAdj)
				}
			}
		} else if color[n] == grey {
			for _, revAdj := range revAdjList[n] {
				dp[n] = (dp[n] + dp[revAdj]) % C
			}
			color[n] = black
		}
	}
	return dp[to]
}

func toReverseAdjacencyList(nodesCount int, edges []edge) [][]int {
	reverseAdjList := make([][]int, nodesCount+1)
	for _, edge := range edges {
		reverseAdjList[edge.to] = append(reverseAdjList[edge.to], edge.from)
	}
	return reverseAdjList
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
	from, to := raw[0], raw[1]

	fmt.Println(countPaths(nodesCount, edges, from, to))
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
