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

var (
	errEmpty = errors.New("queue empty")
)

type node struct {
	value *int
	next  *node
	prev  *node
}

type queue struct {
	tail *node
	head *node
	size int
}

func (q *queue) put(i int) {
	if q.size == 0 {
		node := node{
			value: &i,
		}
		q.head = &node
		q.tail = &node
		q.size++
		return
	}

	q.tail.next = &node{
		value: &i,
		prev:  q.tail,
	}
	q.tail = q.tail.next
	q.size++
}

func (q *queue) get() (int, error) {
	if q.size == 0 {
		return 0, errEmpty
	}
	item := q.head.value
	q.head = q.head.next
	q.size--
	return *item, nil
}

func eval(cmd string) {
	parts := strings.Split(cmd, " ")
	switch parts[0] {
	case "put":
		i, _ := strconv.Atoi(parts[1])
		Q.put(i)
	case "get":
		i, err := Q.get()
		if err != nil {
			fmt.Println("error")
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
	Q = queue{}
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
