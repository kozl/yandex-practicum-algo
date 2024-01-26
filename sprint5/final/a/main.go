package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type heap struct {
	data []person
}

func newHeap() heap {
	return heap{
		data: make([]person, 1),
	}
}

func (h *heap) Add(p person) {
	idx := len(h.data)
	h.data = append(h.data, p)
	h.siftUp(idx)
}

func (h *heap) siftUp(idx int) {
	if idx == 1 {
		return
	}
	if less(h.data[parent(idx)], h.data[idx]) {
		h.data[parent(idx)], h.data[idx] = h.data[idx], h.data[parent(idx)]
		h.siftUp(parent(idx))
	}
}

func (h *heap) siftDown(idx int) {
	if idx >= len(h.data) {
		return
	}
	maxChild, err := h.maxChild(idx)
	if err != nil {
		return
	}

	if less(h.data[idx], h.data[maxChild]) {
		h.data[idx], h.data[maxChild] = h.data[maxChild], h.data[idx]
		h.siftDown(maxChild)
	}

}

var noChildrenErr = errors.New("no children")

func (h *heap) maxChild(idx int) (int, error) {
	left, right := children(idx)
	if right >= len(h.data) && left >= len(h.data) {
		return 0, noChildrenErr
	}
	if right >= len(h.data) {
		return left, nil
	}
	if left >= len(h.data) {
		return right, nil
	}
	maxC := left
	if less(h.data[maxC], h.data[right]) {
		maxC = right
	}
	return maxC, nil
}

func (h *heap) Get() person {
	max := h.data[1]
	h.data[1], h.data[len(h.data)-1] = h.data[len(h.data)-1], h.data[1]
	h.data = h.data[:len(h.data)-1]
	h.siftDown(1)
	return max
}

func sort(array []person) []person {
	h := newHeap()
	for _, person := range array {
		h.Add(person)
	}
	result := make([]person, len(array))
	for i := 0; i < len(array); i++ {
		result[i] = h.Get()
	}
	return result
}

func less(a, b person) bool {
	if a.solved != b.solved {
		return a.solved < b.solved
	}
	if a.penalty != b.penalty {
		return a.penalty > b.penalty
	}
	return a.login > b.login
}

func children(idx int) (left, right int) {
	return idx * 2, idx*2 + 1
}

func parent(idx int) int {
	return idx / 2
}

type person struct {
	login   string
	solved  int
	penalty int
}

func main() {
	scanner := makeScanner()
	num := readInt(scanner)
	persons := make([]person, num)
	for i := 0; i < num; i++ {
		s := readString(scanner)
		persons[i] = personFromString(s)
	}
	sorted := sort(persons)
	for _, person := range sorted {
		fmt.Println(person.login)
	}
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 3 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func readInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	stringInt := scanner.Text()
	res, _ := strconv.Atoi(stringInt)
	return res
}

func readString(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
}

func personFromString(s string) person {
	parts := strings.Split(s, " ")
	login := parts[0]
	solved, err := strconv.Atoi(parts[1])
	if err != nil {
		panic(err.Error())
	}
	penalty, err := strconv.Atoi(parts[2])
	if err != nil {
		panic(err.Error())
	}

	return person{
		login:   login,
		solved:  solved,
		penalty: penalty,
	}
}
