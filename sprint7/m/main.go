package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solve(m int, items []item) []int {
	n := len(items)
	dp := makeDP(n, m)
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if j-items[i-1].weight < 0 {
				dp[i][j] = dp[i-1][j]
				continue
			}
			dp[i][j] = max(dp[i-1][j], items[i-1].cost+dp[i-1][j-items[i-1].weight])
		}
	}

	var result []int
	i, j := n, m
	for j > 0 && i > 0 {
		if dp[i][j] > dp[i-1][j] {
			result = append(result, i)
			j -= items[i-1].weight
			i--
		} else {
			i--
		}
	}

	return result
}

func makeDP(n, m int) [][]int {
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
	}
	return dp
}

type item struct {
	weight int
	cost   int
}

func main() {
	scanner := makeScanner()
	raw := readArray(scanner)
	n, m := raw[0], raw[1]
	var items []item
	for i := 0; i < n; i++ {
		raw := readArray(scanner)
		items = append(items, item{
			weight: raw[0],
			cost:   raw[1],
		})
	}

	result := solve(m, items)
	fmt.Println(len(result))
	printArray(result)
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
