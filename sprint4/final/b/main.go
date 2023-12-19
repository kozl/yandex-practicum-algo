package main

import (
	"bufio"
	"container/list"
	"errors"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

const a float64 = 0.6180339887

type hashMap struct {
	data []*list.List
}

type item struct {
	key   int
	value int
}

func newMap(buckets int) *hashMap {
	return &hashMap{
		data: make([]*list.List, buckets),
	}
}

var errkeyNotFound = errors.New("key not found")

func (m *hashMap) Get(key int) (int, error) {
	b := m.bucket(key)

	el, err := m.findElementByKey(key, m.data[b])
	if err != nil {
		return 0, err
	}
	return el.Value.(item).value, nil
}

func (m *hashMap) Put(key, value int) {
	b := m.bucket(key)
	if m.data[b] == nil {
		l := list.New()
		l.PushFront(item{key: key, value: value})
		m.data[b] = l
		return
	}

	el, err := m.findElementByKey(key, m.data[b])
	if err != nil {
		l := m.data[b]
		l.PushFront(item{key: key, value: value})
		return
	}

	item := el.Value.(item)
	item.value = value
	el.Value = item
}

func (m *hashMap) Delete(key int) (int, error) {
	b := m.bucket(key)
	el, err := m.findElementByKey(key, m.data[b])
	if err != nil {
		return 0, err
	}
	item := el.Value.(item)
	m.data[b].Remove(el)

	return item.value, nil
}

func (m *hashMap) bucket(key int) int {
	_, frac := math.Modf(float64(key) * a)
	in, _ := math.Modf(float64(len(m.data)) * math.Abs(frac))
	return int(in)
}

func (m *hashMap) findElementByKey(key int, l *list.List) (*list.Element, error) {
	if l == nil {
		return nil, errkeyNotFound
	}
	el := l.Front()
	if el == nil {
		return nil, errkeyNotFound
	}
	i := el.Value.(item)
	if i.key == key {
		return el, nil
	}

	for el.Next() != nil {
		el = el.Next()
		i := el.Value.(item)
		if i.key == key {
			return el, nil
		}
	}

	return nil, errkeyNotFound
}

func eval(m *hashMap, cmd string, w io.Writer) {
	parts := strings.Split(cmd, " ")
	switch parts[0] {
	case "get":
		key, _ := strconv.Atoi(parts[1])
		value, err := m.Get(key)
		if errors.Is(err, errkeyNotFound) {
			fmt.Fprintln(w, "None")
			break
		}
		fmt.Fprintln(w, value)
	case "put":
		key, _ := strconv.Atoi(parts[1])
		value, _ := strconv.Atoi(parts[2])
		m.Put(key, value)
	case "delete":
		key, _ := strconv.Atoi(parts[1])
		value, err := m.Delete(key)
		if errors.Is(err, errkeyNotFound) {
			fmt.Fprintln(w, "None")
			break
		}
		fmt.Fprintln(w, value)
	}
}

func main() {
	scanner := makeScanner()
	lines := readInt(scanner)
	m := newMap(1<<22 - 1)
	bufWriter := bufio.NewWriterSize(os.Stdout, 1<<16)
	for i := 0; i < lines; i++ {
		command := readLine(scanner)
		eval(m, command, bufWriter)
	}
	bufWriter.Flush()
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
