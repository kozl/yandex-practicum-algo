package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func solve(total int, denominaitons []int) (banknotesCount int) {
	dp := make([]int, total+1)
	for i := 0; i <= total; i++ {
		dp[i] = math.MaxInt32
	}

	for i := 1; i <= len(denominaitons); i++ {
		for j := 1; j <= total; j++ {
			if j-denominaitons[i-1] > 0 {
				dp[j] = min(dp[j-denominaitons[i-1]]+1, dp[j])
				continue
			}
			if j-denominaitons[i-1] == 0 {
				dp[j] = 1
			}
		}
	}
	return dp[total]
}

func main() {
	scanner := makeScanner()
	total := readInt(scanner)
	_ = readInt(scanner)
	denominaitons := readArray(scanner)
	sanitize(&denominaitons)
	banknotesCount := solve(total, denominaitons)
	fmt.Println(banknotesCount)
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

func sanitize(denominations *[]int) {
	m := map[int]bool{}
	for _, i := range *denominations {
		m[i] = true
	}
	result := []int{}
	for k := range m {
		result = append(result, k)
	}
	sort.Ints(result)
	*denominations = result
}
