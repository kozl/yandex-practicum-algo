package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(strs []string) int {
	var prefixlen int
	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if strs[0][i] != strs[j][i] {
				return prefixlen
			}
		}
		prefixlen++
	}
	return prefixlen
}

func main() {
	scanner := makeScanner()
	count := readInt(scanner)
	strs := []string{}
	for i := 0; i < count; i++ {
		strs = append(strs, readString(scanner))
	}
	fmt.Println(solve(strs))
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
