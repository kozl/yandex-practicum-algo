package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func solve(line, pattern, replacement string) string {
	s := pattern + "#" + line
	prefixFun := make([]int, len(s))
	positions := []int{}
	for i := 1; i < len(s); i++ {
		k := prefixFun[i-1]
		for k > 0 && s[i] != s[k] {
			k = prefixFun[k-1]
		}
		if s[i] == s[k] {
			k++
		}
		prefixFun[i] = k
		if k == len(pattern) {
			positions = append(positions, i-2*len(pattern))
		}
	}
	result := strings.Builder{}
	var start int
	for _, pos := range positions {
		result.WriteString(line[start:pos])
		result.WriteString(replacement)
		start = pos + len(pattern)
	}
	if start < len(line) {
		result.WriteString(line[start:])
	}
	return result.String()
}

func main() {
	scanner := makeScanner()
	line := readString(scanner)
	pattern := readString(scanner)
	replacement := readString(scanner)
	fmt.Println(solve(line, pattern, replacement))
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 20 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func readString(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
}

func printArray(w io.Writer, array []int) {
	s := make([]string, len(array))
	for i := 0; i < len(array); i++ {
		s[i] = strconv.Itoa(array[i])
	}
	fmt.Fprintln(w, strings.Join(s, " "))
}
