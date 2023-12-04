package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func polyHash(s string, a, mod int) int {
	hash := 0
	b := 1
	for i := 0; i < len(s); i++ {
		hash += int(s[len(s)-1-i]) * b
		b = (b * a) % mod
		hash = hash % mod
	}
	return hash
}

func main() {
	scanner := makeScanner()
	a := readInt(scanner)
	mod := readInt(scanner)
	s := readString(scanner)
	result := polyHash(s, a, mod)
	fmt.Println(result)
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
