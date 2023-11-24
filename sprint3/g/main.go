package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func countSort(array []int) {
	countValues := make([]int, 3)
	for _, arr := range array {
		countValues[arr]++
	}

	pos := 0
	for i := 0; i < len(array); i++ {
		for countValues[pos] == 0 {
			pos++
		}
		array[i] = pos
		countValues[pos]--
	}
}

func main() {
	scanner := makeScanner()
	_ = readInt(scanner)
	array := readArray(scanner)
	countSort(array)
	printArray(array)
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
	if len(listString) == 1 && listString[0] == "" {
		return []int{}
	}
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

func printArray(array []int) {
	if len(array) == 0 {
		return
	}
	s := make([]string, len(array))
	for idx, i := range array {
		s[idx] = strconv.Itoa(i)
	}
	fmt.Println(strings.Join(s, " "))
}
