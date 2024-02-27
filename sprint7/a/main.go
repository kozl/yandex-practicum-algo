package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solve(days int, prices []int) int {
	monotonicIncr := [][2]int{}

	prev := prices[0]
	seqStart := 0
	seqEnd := 0
	for idx := 1; idx < len(prices); idx++ {
		if prices[idx] >= prev {
			seqEnd = idx
		} else {
			if seqEnd-seqStart >= 1 {
				monotonicIncr = append(monotonicIncr, [2]int{seqStart, seqEnd})
			}
			seqStart = idx
			seqEnd = idx
		}
		prev = prices[idx]
	}

	if seqEnd-seqStart >= 1 {
		monotonicIncr = append(monotonicIncr, [2]int{seqStart, seqEnd})
	}

	sum := 0
	for _, seq := range monotonicIncr {
		sum += prices[seq[1]] - prices[seq[0]]
	}
	return sum
}

func main() {
	scanner := makeScanner()
	days := readInt(scanner)
	prices := readArray(scanner)
	fmt.Println(solve(days, prices))
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
