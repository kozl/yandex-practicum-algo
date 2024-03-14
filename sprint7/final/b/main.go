package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solve(scores []int) bool {
	s := sum(scores)
	if s%2 == 1 {
		return false
	}
	dp := makeDP(len(scores), s/2)

	for i := 1; i <= len(scores); i++ {
		for j := 1; j <= s/2; j++ {
			dp[i][j] = get(dp, i-1, j) || get(dp, i-1, j-scores[i-1])
		}
	}

	return dp[len(scores)][s/2]
}

func get(dp [][]bool, i, j int) bool {
	if i < 0 || i > len(dp) {
		return false
	}
	if j < 0 || j > len(dp[0]) {
		return false
	}
	return dp[i][j]
}

func sum(arr []int) (sum int) {
	for _, i := range arr {
		sum += i
	}
	return sum
}

func makeDP(n, m int) [][]bool {
	dp := make([][]bool, n+1)
	for idx := range dp {
		dp[idx] = make([]bool, m+1)
		dp[idx][0] = true
	}
	return dp
}

func main() {
	scanner := makeScanner()
	_ = readInt(scanner)
	scores := readArray(scanner)
	printBool(solve(scores))
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

func printBool(b bool) {
	if b {
		fmt.Println("True")
		return
	}
	fmt.Println("False")
}
