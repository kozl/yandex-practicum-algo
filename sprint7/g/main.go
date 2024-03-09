package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solve(m int, coins []int) int {
	n := len(coins)
	dp := makeDP(n, m)
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if j-coins[i-1] < 0 {
				dp[i][j] = dp[i-1][j]
				continue
			}
			if j-coins[i-1] == 0 {
				dp[i][j] = dp[i-1][j] + 1
				continue
			}
			dp[i][j] = dp[i-1][j] + dp[i][j-coins[i-1]]
		}
	}
	return dp[n][m]
}

func makeDP(n, m int) [][]int {
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
	}
	return dp
}

func main() {
	scanner := makeScanner()
	m := readInt(scanner)
	_ = readInt(scanner)
	coins := readArray(scanner)
	fmt.Println(solve(m, coins))
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
