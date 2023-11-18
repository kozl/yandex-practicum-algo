package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type stack struct {
	data [][2]int
}

var S = stack{
	data: [][2]int{},
}

func (s *stack) push(i int) {
	var max int
	if len(s.data) != 0 {
		max = s.data[len(s.data)-1][1]
		if i > max {
			max = i
		}
	} else {
		max = i
	}
	s.data = append(s.data, [2]int{i, max})
}

var emptyErr = errors.New("stack is empty")

func (s *stack) pop() (int, error) {
	if len(s.data) == 0 {
		return 0, emptyErr
	}
	lastIdx := len(s.data) - 1
	item := s.data[lastIdx]
	s.data = s.data[:lastIdx]
	return item[0], nil
}

func (s *stack) getMax() (int, error) {
	if len(s.data) == 0 {
		return 0, emptyErr
	}
	return s.data[len(s.data)-1][1], nil
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
