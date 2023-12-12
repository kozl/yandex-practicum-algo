package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	traslation := map[rune]rune{}
	traslationReverse := map[rune]rune{}
	for i := 0; i < len(a); i++ {
		if r, ok := traslation[rune(a[i])]; ok {
			if rune(b[i]) != r {
				return false
			}
		} else {
			if _, ok := traslationReverse[rune(b[i])]; ok {
				return false
			}
			traslation[rune(a[i])] = rune(b[i])
			traslationReverse[rune(b[i])] = rune(a[i])
		}
	}
	return true
}

func main() {
	scanner := makeScanner()
	a := readString(scanner)
	b := readString(scanner)
	result := solve(a, b)
	if result {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
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
