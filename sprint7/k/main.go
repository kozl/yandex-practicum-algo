package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func solve(a, b []int) (s int, aIdx, bIdx []int) {
	dp := make([][]int, len(a)+1)
	for idx := range dp {
		dp[idx] = make([]int, len(b)+1)
	}
	for i := 1; i <= len(a); i++ {
		for j := 1; j <= len(b); j++ {
			if a[i-1] == b[j-1] {
				dp[i][j] = get(dp, i-1, j-1) + 1
			} else {
				dp[i][j] = max(get(dp, i-1, j), get(dp, i, j-1))
			}
		}
	}

	s = dp[len(a)][len(b)]
	i, j := len(a), len(b)

	for i != 0 && j != 0 {
		if a[i-1] == b[j-1] {
			aIdx = append(aIdx, i)
			bIdx = append(bIdx, j)
			i--
			j--
		} else {
			if dp[i][j] == dp[i-1][j] {
				i--
			} else {
				j--
			}
		}
	}
	slices.Reverse(aIdx)
	slices.Reverse(bIdx)

	return s, aIdx, bIdx
}

func get(dp [][]int, i, j int) int {
	if i < 0 || j < 0 || i >= len(dp) || j >= len(dp[0]) {
		return 0
	}
	return dp[i][j]
}

func main() {
	scanner := makeScanner()
	_ = readInt(scanner)
	a := readArray(scanner)
	_ = readInt(scanner)
	b := readArray(scanner)
	c, aIdx, bIdx := solve(a, b)
	fmt.Println(c)
	printArray(aIdx)
	printArray(bIdx)
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

func printArray(arr []int) {
	s := []string{}
	for _, i := range arr {
		s = append(s, fmt.Sprint(i))
	}
	fmt.Println(strings.Join(s, " "))
}
