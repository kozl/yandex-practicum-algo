package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func solve(s string) string {
	lc := letterCount(s)
	var sb strings.Builder
	for idx := range lc {
		if lc[idx].count > 0 {
			for lc[idx].count > 1 {
				sb.WriteRune(lc[idx].letter)
				lc[idx].count -= 2
			}
		}
	}
	var middleRune rune
	for _, item := range lc {
		if item.count == 1 {
			middleRune = item.letter
			break
		}
	}
	if middleRune != 0 {
		return sb.String() + string(middleRune) + reverse(sb.String())
	}
	return sb.String() + reverse(sb.String())
}

func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

type item struct {
	letter rune
	count  int
}

func letterCount(s string) []item {
	var result []item
	m := make(map[rune]int)
	for _, r := range s {
		m[r]++
	}
	for k, v := range m {
		result = append(result, item{letter: k, count: v})
	}
	sort.Slice(result, func(i, j int) bool { return result[i].letter < result[j].letter })
	return result
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
