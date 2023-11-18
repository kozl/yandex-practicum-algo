package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var Q queue

func newQueue(qSize int) queue {
	return queue{
		data: make([]int, qSize),
		tail: 0,
		head: 0,
		maxN: qSize,
		size: 0,
	}
}

var (
	errFull  = errors.New("queue full")
	errEmpty = errors.New("queue empty")
)

type queue struct {
	data []int
	tail int
	head int
	maxN int
	size int
}

func (q *queue) push(i int) error {
	if q.size == q.maxN {
		return errFull
	}
	q.data[q.tail] = i
	q.tail = (q.tail + 1) % q.maxN
	q.size++
	return nil
}

func (q *queue) pop() (int, error) {
	if q.size == 0 {
		return 0, errEmpty
	}
	item := q.data[q.head]
	q.head = (q.head + 1) % q.maxN
	q.size--
	return item, nil
}

func (q *queue) peek() (int, error) {
	if q.size == 0 {
		return 0, errEmpty
	}
	return q.data[q.head], nil
}

func eval(cmd string) {
	parts := strings.Split(cmd, " ")
	switch parts[0] {
	case "push":
		i, _ := strconv.Atoi(parts[1])
		err := Q.push(i)
		if err != nil {
			fmt.Println("error")
		}
	case "pop":
		i, err := Q.pop()
		if err != nil {
			fmt.Println("None")
			break
		}
		fmt.Println(i)
	case "peek":
		i, err := Q.peek()
		if err != nil {
			fmt.Println("None")
			break
		}
		fmt.Println(i)
	case "size":
		fmt.Println(Q.size)
	}
}

func main() {
	scanner := makeScanner()
	lines := readInt(scanner)
	qsize := readInt(scanner)
	Q = newQueue(qsize)
	for i := 0; i < lines; i++ {
		command := readLine(scanner)
		eval(command)
	}
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
