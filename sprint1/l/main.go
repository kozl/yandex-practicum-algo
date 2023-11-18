package main

import (
	"bufio"
	"fmt"
	"os"
)

func getExcessiveLetter(s1 string, t1 string) string {
	runeset := map[rune]int{}
	for _, r := range t1 {
		runeset[r]++
	}

	for _, r := range s1 {
		runeset[r]--
	}

	for k, v := range runeset {
		if v == 1 {
			return string(k)
		}
	}
	return ""
}

func main() {
	scanner := makeScanner()
	s := readLine(scanner)
	t := readLine(scanner)
	fmt.Print(getExcessiveLetter(s, t))
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
