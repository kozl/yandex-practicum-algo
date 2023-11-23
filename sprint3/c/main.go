package main

import (
	"bufio"
	"fmt"
	"os"
)

func isSubsequence(s, t string) bool {
	if len(s) > len(t) {
		return false
	}
	var i, j int
	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			i++
		}
		j++
	}
	return i == len(s)
}

func main() {
	scanner := makeScanner()
	s := readLine(scanner)
	t := readLine(scanner)
	result := isSubsequence(s, t)
	if result {
		fmt.Println("True")
		return
	}
	fmt.Println("False")
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
