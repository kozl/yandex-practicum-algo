package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var C = int(math.Pow(float64(10), float64(9))) + 7

func solve(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	dp1, dp2 := 1, 1
	var result int
	for i := 2; i <= n; i++ {
		result = (dp1 + dp2) % C
		dp1 = dp2
		dp2 = result
	}
	return result
}

func main() {
	scanner := makeScanner()
	n := readInt(scanner)
	fmt.Println(solve(n))
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 3 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func readInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	stringInt := scanner.Text()
	res, _ := strconv.Atoi(stringInt)
	return res
}
