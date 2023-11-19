package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var (
	errFull  = errors.New("dequeue full")
	errEmpty = errors.New("dequeue empty")
)

type dequeue struct {
	data []int
	head int
	size int
	tail int
}

func newDeq(size int) *dequeue {
	return &dequeue{
		data: make([]int, size),
		head: 0,
		size: 0,
		tail: 0,
	}
}

func (d *dequeue) pushBack(i int) error {
	if d.size >= cap(d.data) {
		return errFull
	}
	d.data[d.tail] = i
	d.tail = (d.tail + 1) % cap(d.data)
	d.size++
	return nil
}

func (d *dequeue) pushFront(i int) error {
	if d.size >= cap(d.data) {
		return errFull
	}
	d.head--
	if d.head < 0 {
		d.head = cap(d.data) - 1
	}
	d.data[d.head] = i
	d.size++
	return nil
}

func (d *dequeue) popFront() (int, error) {
	if d.size == 0 {
		return 0, errEmpty
	}
	elem := d.data[d.head]
	d.head = (d.head + 1) % cap(d.data)
	d.size--
	return elem, nil
}

func (d *dequeue) popBack() (int, error) {
	if d.size == 0 {
		return 0, errEmpty
	}
	d.tail--
	if d.tail < 0 {
		d.tail = cap(d.data) - 1
	}
	d.size--
	return d.data[d.tail], nil
}

func eval(d *dequeue, cmd string, w io.Writer) {
	parts := strings.Split(cmd, " ")
	switch parts[0] {
	case "push_back":
		i, _ := strconv.Atoi(parts[1])
		if err := d.pushBack(i); err != nil {
			fmt.Fprintln(w, "error")
		}
	case "push_front":
		i, _ := strconv.Atoi(parts[1])
		if err := d.pushFront(i); err != nil {
			fmt.Fprintln(w, "error")
		}
	case "pop_front":
		value, err := d.popFront()
		if err != nil {
			fmt.Fprintln(w, "error")
			break
		}
		fmt.Fprintln(w, value)
	case "pop_back":
		value, err := d.popBack()
		if err != nil {
			fmt.Fprintln(w, "error")
			break
		}
		fmt.Fprintln(w, value)
	}
}

func main() {
	scanner := makeScanner()
	commands := readInt(scanner)
	deqSize := readInt(scanner)
	D := newDeq(deqSize)
	for i := 0; i < commands; i++ {
		command := readLine(scanner)
		eval(D, command, os.Stdout)
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
