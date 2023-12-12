package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solve(array []int) int {
	var max int
	var allSum int

	resultIdx := make(map[int][]int)
	resultIdx[0] = append(resultIdx[0], 0)

	for i := 0; i < len(array); i++ {
		if array[i] == 0 {
			allSum++
			resultIdx[allSum] = append(resultIdx[allSum], i+1)
		} else {
			allSum--
			resultIdx[allSum] = append(resultIdx[allSum], i+1)
		}
	}
	for _, v := range resultIdx {
		if len(v) == 0 {
			continue
		}
		curr := v[len(v)-1] - v[0]
		if curr > max {
			max = curr
		}
	}

	return max
}

func main() {
	scanner := makeScanner()
	_ = readInt(scanner)
	array := readArray(scanner)
	result := solve(array)
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

func readInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	stringInt := scanner.Text()
	res, _ := strconv.Atoi(stringInt)
	return res
}
