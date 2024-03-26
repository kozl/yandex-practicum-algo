package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func solve(s, t string) int {
	var sb, tb strings.Builder
	for i := 0; i < len(s); i++ {
		if pos(s[i])%2 == 0 {
			sb.WriteByte(s[i])
		}
	}

	for i := 0; i < len(t); i++ {
		if pos(t[i])%2 == 0 {
			tb.WriteByte(t[i])
		}
	}
	return strings.Compare(sb.String(), tb.String())
}

func pos(b byte) int {
	return int(rune(b)-'a') + 1
}

func main() {
	scanner := makeScanner()
	s := readString(scanner)
	t := readString(scanner)
	fmt.Println(solve(s, t))
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
