package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func solve(s string, operations []operation) string {
	sort.Slice(operations, func(i, j int) bool { return operations[i].pos < operations[j].pos })
	var pos int
	var result strings.Builder
	for _, op := range operations {
		for i := pos; i < op.pos; i++ {
			result.WriteByte(s[i])
		}
		result.WriteString(op.s)
		pos = op.pos
	}
	for i := pos; i < len(s); i++ {
		result.WriteByte(s[i])
	}
	return result.String()
}

type operation struct {
	s   string
	pos int
}

func main() {
	scanner := makeScanner()
	s := readString(scanner)
	inserts := readInt(scanner)
	var operations []operation
	for i := 0; i < inserts; i++ {
		raw := readString(scanner)
		parts := strings.Split(raw, " ")
		t := parts[0]
		k, _ := strconv.Atoi(parts[1])
		operations = append(operations, operation{
			s:   t,
			pos: k,
		})
	}
	fmt.Println(solve(s, operations))
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 3 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func readString(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
}

func readInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	stringInt := scanner.Text()
	res, _ := strconv.Atoi(stringInt)
	return res
}
