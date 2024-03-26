package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func solve(s, t string) bool {
	if int(math.Abs(float64(len(s)-len(t)))) > 1 {
		return false
	}
	if len(s) == len(t) {
		for i := 0; i < len(s); i++ {
			if s[i] == t[i] {
				continue
			}
			newStr := s[:i] + string(t[i]) + s[i+1:]
			return newStr == t
		}
		return true
	}
	var short, long string
	if len(s) > len(t) {
		short = t
		long = s
	} else {
		short = s
		long = t
	}

	for i := 0; i < len(short); i++ {
		if short[i] == long[i] {
			continue
		}
		newStr := long[:i] + long[i+1:]
		return short == newStr
	}
	return true
}

func main() {
	scanner := makeScanner()
	s := readString(scanner)
	t := readString(scanner)
	printBool(solve(s, t))
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

func printBool(b bool) {
	if b {
		fmt.Println("OK")
		return
	}
	fmt.Println("FAIL")
}
