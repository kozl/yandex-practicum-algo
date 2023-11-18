package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func fibMod(n, k int) int {
	if n == 1 || n == 0 {
		return 1
	}
	mod := int(math.Pow(10, float64(k)))
	a, b := 1, 1

	for i := 2; i <= n; i++ {
		a = (a + b) % mod
		a, b = b, a
	}
	return b
}

func main() {
	scanner := makeScanner()
	args := readArray(scanner)
	result := fibMod(args[0], args[1])
	fmt.Println(result)
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
