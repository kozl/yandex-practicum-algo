package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve(s string) int {
	set := map[rune]bool{}
	i, j, l := 0, 0, 0
	for j < len(s) {
		if !set[rune(s[j])] {
			set[rune(s[j])] = true
			j++
			l = max(l, len(set))
		} else {
			i++
			set[rune(s[i])] = false
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
