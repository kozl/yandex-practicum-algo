package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func generateSequences(s string, open int, close int, n int, results *[]string) {
	if len(s) == 2*n {
		*results = append(*results, s)
		return
	}
	if open < n {
		generateSequences(s+"(", open+1, close, n, results)
	}
	if close < open {
		generateSequences(s+")", open, close+1, n, results)
	}
}

func result(n int) []string {
	result := []string{}
	generateSequences("", 0, 0, n, &result)
	return result
}

func main() {
	scanner := makeScanner()
	count := readInt(scanner)
	bracketSequences := result(count)

	sort.Strings(bracketSequences)
	for _, seq := range bracketSequences {
		fmt.Println(seq)
	}
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 20 * 1024 * 1024
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
