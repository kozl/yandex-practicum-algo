package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func solve(sentence string) string {
	words := strings.Split(sentence, " ")
	if len(words) <= 1 {
		return sentence
	}
	for i := 0; i < len(words)/2; i++ {
		words[i], words[len(words)-i-1] = words[len(words)-i-1], words[i]
	}
	return strings.Join(words, " ")
}

func main() {
	scanner := makeScanner()
	sentence := readString(scanner)
	fmt.Println(solve(sentence))
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
