package main

import (
	"bufio"
	"os"
	"slices"
	"strconv"
)

func getBinaryNumber(n int) []int {
	var q, r int
	q = n / 2
	r = n % 2
	result := []int{}
	for q != 0 {
		result = append(result, r)
		r = q % 2
		q = q / 2
	}
	result = append(result, r)
	slices.Reverse(result)
	return result
}

func main() {
	scanner := makeScanner()
	n := readInt(scanner)
	binaryNumber := getBinaryNumber(n)
	printArray(binaryNumber)
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 3 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func readLine(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
}

func readInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	stringInt := scanner.Text()
	res, _ := strconv.Atoi(stringInt)
	return res
}

func printArray(arr []int) {
	writer := bufio.NewWriter(os.Stdout)
	for i := 0; i < len(arr); i++ {
		writer.WriteString(strconv.Itoa(arr[i]))
	}
	writer.Flush()
}
