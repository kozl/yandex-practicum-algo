package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func solve(houses []int, budget int) int {
	sort.Ints(houses)

	result := 0
	for _, housePrice := range houses {
		if budget-housePrice >= 0 {
			budget = budget - housePrice
			result++
		} else {
			break
		}
	}
	return result
}

func main() {
	scanner := makeScanner()
	arr1 := readArray(scanner)
	houses := readArray(scanner)
	fmt.Println(solve(houses, arr1[1]))
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 20 * 1024 * 1024
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
