package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func solve(a []int) (indices []int) {
	prev := make([]int, len(a))
	dp := make([]int, len(a))
	for i := 0; i < len(a); i++ {
		dp[i] = 1
		prev[i] = -1
		for j := 0; j < i; j++ {
			if (a[j] < a[i]) && (dp[j]+1 > dp[i]) {
				dp[i] = dp[j] + 1
				prev[i] = j
			}
		}
	}

	pos := 0
	length := dp[0]
	for i := 0; i < len(a); i++ {
		if dp[i] > length {
			pos = i
			length = dp[i]
		}
	}

	for pos != -1 {
		indices = append(indices, pos+1)
		pos = prev[pos]
	}
	slices.Reverse(indices)
	return indices
}

func main() {
	scanner := makeScanner()
	_ = readInt(scanner)
	a := readArray(scanner)
	indices := solve(a)
	fmt.Println(len(indices))
	printArray(indices)
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
