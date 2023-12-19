package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve(s string) int {
	set := map[string]bool{}
	i, j, l := 0, 0, 0
	for j < len(s) {
		if !set[string(s[j])] {
			set[string(s[j])] = true
			j++
			l = max(l, len(set))
		} else {
			delete(set, string(s[i]))
			i++
		}
	}
	return l
}

func main() {
	scanner := makeScanner()
	s := readString(scanner)
	fmt.Println(solve(s))
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
