package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	mod = 1000000000
	A   = 131
)

func solve(a, b []int) int {
	left := 0
	right := min(len(a), len(b))
	maxN := 0
	for left <= right {
		cur := (left + right) / 2
		if hasCommonArrayLen(a, b, cur) {
			maxN = max(maxN, cur)
			left = cur + 1
		} else {
			right = cur - 1
		}
	}
	return maxN
}

func hasCommonArrayLen(a, b []int, n int) bool {
	if n > len(b) || n > len(a) {
		return false
	}
	for i := 0; i+n <= len(a); i++ {
		for j := 0; j+n <= len(b); j++ {
			h1 := hash(a[i:i+n], A, mod)
			h2 := hash(b[j:j+n], A, mod)
			if h1 == h2 {
				return true
			}
		}
	}
	return false
}

func hash(arr []int, n, mod int) int {
	result := 0
	N := 1
	for _, v := range arr {
		result += (v * N) % mod
		N = (N * n) % mod
	}
	return result
}

func main() {
	scanner := makeScanner()
	_ = readInt(scanner)
	a := readArray(scanner)
	_ = readInt(scanner)
	b := readArray(scanner)
	fmt.Println(solve(a, b))
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 3 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func readArray(scanner *bufio.Scanner) []int {
	scanner.Scan()
	listString := strings.Split(scanner.Text(), " ")
	arr := make([]int, len(listString))
	for i := 0; i < len(listString); i++ {
		arr[i], _ = strconv.Atoi(listString[i])
	}
	return arr
}

func readInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	stringInt := scanner.Text()
	res, _ := strconv.Atoi(stringInt)
	return res
}
