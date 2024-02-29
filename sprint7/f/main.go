package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var modC = int(math.Pow(float64(10), float64(9))) + 7

func solve(n int, k int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	for i := 2; i <= n; i++ {
		for j := previous(i, k); j < i; j++ {
			dp[i] += dp[j]
			dp[i] = dp[i] % modC
		}
	}
	return dp[n]
}

func previous(i, k int) int {
	if i-k <= 0 {
		return 0
	}
	return i - k
}

func main() {
	scanner := makeScanner()
	raw := readArray(scanner)
	n, k := raw[0], raw[1]
	fmt.Println(solve(n, k))
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
