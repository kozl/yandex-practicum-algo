package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type stack struct {
	data []int
}

var S = stack{
	data: []int{},
}

func (s *stack) push(i int) {
	s.data = append(s.data, i)
}

var emptyErr = errors.New("stack is empty")

func (s *stack) pop() (int, error) {
	if len(s.data) == 0 {
		return 0, emptyErr
	}
	lastIdx := len(s.data) - 1
	item := s.data[lastIdx]
	s.data = s.data[:lastIdx]
	return item, nil
}

func (s *stack) getMax() (int, error) {
	if len(s.data) == 0 {
		return 0, emptyErr
	}
	return slices.Max(s.data), nil
}

func eval(cmd string) {
	parts := strings.Split(cmd, " ")
	switch parts[0] {
	case "push":
		i, _ := strconv.Atoi(parts[1])
		S.push(i)
	case "pop":
		_, err := S.pop()
		if err != nil {
			fmt.Println("error")
		}
	case "get_max":
		i, err := S.getMax()
		if err != nil {
			fmt.Println("None")
			break
		}
		fmt.Println(i)
	}
}

func main() {
	scanner := makeScanner()
	lines := readInt(scanner)
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
