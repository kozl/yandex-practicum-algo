package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solve(n, m int, gold []int) (maxWeight int) {
	current, previous := make([]int, m+1), make([]int, m+1)
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if getWeight(gold, i) > j {
				current[j] = previous[j]
			} else {
				current[j] = max(previous[j-getWeight(gold, i)]+getWeight(gold, i), previous[j])
			}
		}
		previous = current
		current = make([]int, m+1)
	}
	return previous[m]
}

func getWeight(gold []int, idx int) int {
	if idx == 0 {
		return 0
	}
	return gold[idx-1]
}

func main() {
	scanner := makeScanner()
	raw := readArray(scanner)
	n, m := raw[0], raw[1]
	gold := readArray(scanner)

	fmt.Println(solve(n, m, gold))
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
