package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(words []string) string {
	m := map[string]int{}
	var (
		max     int
		maxWord string
	)

	for _, word := range words {
		m[word] += 1
		if m[word] > max {
			max = m[word]
			maxWord = word
		}
		if m[word] == max {
			if word < maxWord {
				maxWord = word
			}
		}
	}
	return maxWord
}

func main() {
	scanner := makeScanner()
	count := readInt(scanner)
	words := []string{}
	for i := 0; i < count; i++ {
		words = append(words, readString(scanner))
	}
	fmt.Println(solve(words))
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
