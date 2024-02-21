package main

/*

[Успешная посылка](https://contest.yandex.ru/contest/25070/run-report/107876320/)

Для решения задачи представим граф таким образом: дороги одного цвета будут ориентированы в одну сторону, а дороги другого цвета — в другую.
Выясняется, что граф дорог будет оптимальным тогда, когда в таком графе не будет циклов.
Поэтому для того чтобы выяснить оптимальность достаточно будет проверить граф на циклы при помощи алгоритма обхода DFS.

*/

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

const (
	white = iota
	grey
	black
)

const (
	typeR = 'R'
	typeB = 'B'
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

func isOptimal(cities int, adjacencyList [][]int) bool {
	color := make([]int, cities+1)
	for cityID := 1; cityID < len(color); cityID++ {
		if color[cityID] != white {
			continue
		}
		if dfsHasLoops(cities, cityID, adjacencyList, color) {
			return false
		}
	}
	return true
}

func dfsHasLoops(nodesCount, startNode int, adjList [][]int, color []int) bool {
	st := stack{}
	st.push(startNode)

	for !st.empty() {
		node := st.pop()
		if color[node] == white {
			color[node] = grey
			st.push(node)
			for _, adjacent := range adjList[node] {
				if color[adjacent] == white {
					st.push(adjacent)
				} else if color[adjacent] == grey {
					return true
				}
			}
		} else if color[node] == grey {
			color[node] = black
		}
	}
	return false
}

func toAdjacencyList(cities int, mapDescription []string) [][]int {
	adjList := make([][]int, cities+1)
	for idx, line := range mapDescription {
		cityID := idx + 1
		for idx, roadType := range line {
			adjCityID := cityID + idx + 1
			if roadType == typeB {
				adjList[cityID] = append(adjList[cityID], adjCityID)
			} else {
				adjList[adjCityID] = append(adjList[adjCityID], cityID)
			}
		}
	}
	return adjList
}

func solve(w io.Writer, cities int, mapDescription []string) {
	adjList := toAdjacencyList(cities, mapDescription)
	yes := isOptimal(cities, adjList)
	fmt.Fprintln(w, toStr(yes))
}

func main() {
	scanner := makeScanner()
	cities := readInt(scanner)
	mapDescription := make([]string, cities-1)
	for i := 0; i < cities-1; i++ {
		mapDescription[i] = readLine(scanner)
	}
	solve(os.Stdout, cities, mapDescription)
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 3 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func readLine(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
}

func readInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	stringInt := scanner.Text()
	res, _ := strconv.Atoi(stringInt)
	return res
}

func toStr(yes bool) string {
	if yes {
		return "YES"
	}
	return "NO"
}
