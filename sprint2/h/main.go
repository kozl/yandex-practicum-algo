package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

type stack struct {
	data []rune
}

var S = stack{
	data: []rune{},
}

func (s *stack) push(r rune) {
	s.data = append(s.data, r)
}

var emptyErr = errors.New("stack is empty")

func (s *stack) pop() (rune, error) {
	if len(s.data) == 0 {
		return 0, emptyErr
	}
	lastIdx := len(s.data) - 1
	item := s.data[lastIdx]
	s.data = s.data[:lastIdx]
	return item, nil
}

func isCorrectBracketSeq(seq string) bool {
	for _, r := range seq {
		switch r {
		case '{', '[', '(':
			S.push(r)
		case '}':
			r, err := S.pop()
			if err != nil {
				return false
			}
			if r != '{' {
				return false
			}
		case ')':
			r, err := S.pop()
			if err != nil {
				return false
			}
			if r != '(' {
				return false
			}
		case ']':
			r, err := S.pop()
			if err != nil {
				return false
			}
			if r != '[' {
				return false
			}
		}
	}

	return len(S.data) == 0
}

func main() {
	scanner := makeScanner()
	line := readLine(scanner)
	result := isCorrectBracketSeq(line)
	if result {
		fmt.Println("True")
	} else {
		fmt.Println("False")
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
