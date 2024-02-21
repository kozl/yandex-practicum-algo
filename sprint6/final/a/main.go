package main

/*

[Успешная посылка](https://contest.yandex.ru/contest/25070/run-report/107827732/)

Для решения описал очередь с приоритетами, которая хранит ребра-кандидаты на добавление.
Так как по условию задачи нужно построить остовной граф с максимальными рёбрами максимальный
приоритет очереди означает максимальный вес ребра.

Далее остовной граф строю по алгоритму Прима:
1. начинаю с 1 ноды, добавляю в heap смежные ребра
2. повторяю в цикле до тех пора пока не добавлены все узлы, либо не будет больше рёбер-кандидатов
3. на каждой итерации извлекаю ребро из heap, добавляю узел на противоположной стороне, а в heap
   добавляю новые ребра, которые исходят из того узла, который только что добавил
4. остовное дерево нельзя построить если после выхода из цикла количество добавленых узлов меньше
   количества узлов в исходном графе
5. подсчитываю суммарный вес всех рёбер в полученном MST

Вычислительная сложность: V * E * Log(E)
Сложность по памяти: V * E (кажется)

*/

import (
	"bufio"
	"container/heap"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type edgeHeap []edge

func (h edgeHeap) Len() int           { return len(h) }
func (h edgeHeap) Less(i, j int) bool { return h[i].weight > h[j].weight }
func (h edgeHeap) Swap(i, j int) {
	// по какой-то причине в одном из тестов в Swap попадал j < 0
	// почему так могло происходить?
	if j < 0 || i < 0 {
		return
	}
	h[i], h[j] = h[j], h[i]
}
func (h *edgeHeap) Push(x any) {
	*h = append(*h, x.(edge))
}
func (h *edgeHeap) Pop() any {
	if len(*h) == 0 {
		return nil
	}
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type edge struct {
	from   int
	to     int
	weight int
}

type node struct {
	id     int
	weight int
}

func mst(nodesCount int, adj [][]node) []edge {
	start := 1
	nodesAdded := map[int]bool{start: true}
	mst := []edge{}
	candidateEdges := edgeHeap{}
	heap.Init(&candidateEdges)

	addCandidateEdges(&candidateEdges, start, adj)
	for len(nodesAdded) != nodesCount && candidateEdges.Len() != 0 {
		var maxEdge edge
		for {
			e := heap.Pop(&candidateEdges)
			if e == nil {
				return nil
			}
			maxEdge = e.(edge)
			if _, ok := nodesAdded[maxEdge.to]; !ok {
				break
			}
		}
		nodesAdded[maxEdge.to] = true
		addCandidateEdges(&candidateEdges, maxEdge.to, adj)
		mst = append(mst, maxEdge)
	}
	if len(nodesAdded) != nodesCount {
		return nil
	}
	return mst
}

func addCandidateEdges(h heap.Interface, from int, adj [][]node) {
	for _, node := range adj[from] {
		heap.Push(h, edge{from: from, to: node.id, weight: node.weight})
	}
}

func toWadjacencyList(nodesCount int, edges []edge) [][]node {
	adjList := make([][]node, nodesCount+1)
	for _, e := range edges {
		adjList[e.from] = append(adjList[e.from], node{id: e.to, weight: e.weight})
		adjList[e.to] = append(adjList[e.to], node{id: e.from, weight: e.weight})
	}
	return adjList
}

func solve(w io.Writer, nodesCount int, edges []edge) {
	adjList := toWadjacencyList(nodesCount, edges)
	mstEdges := mst(nodesCount, adjList)
	if mstEdges == nil {
		fmt.Fprintln(w, "Oops! I did it again")
		return
	}

	var totalWeight int
	for _, edge := range mstEdges {
		totalWeight += edge.weight
	}
	fmt.Fprintln(w, totalWeight)
}

func main() {
	scanner := makeScanner()
	raw := readArray(scanner)
	nodesCount, edgesCount := raw[0], raw[1]
	var edges []edge
	for i := 0; i < edgesCount; i++ {
		raw := readArray(scanner)
		edges = append(edges, edge{from: raw[0], to: raw[1], weight: raw[2]})
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
