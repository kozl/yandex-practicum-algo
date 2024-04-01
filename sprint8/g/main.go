package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func find(x, a []int, start int) int {
	if len(a) > len(x[start:]) {
		return -1
	}
	for pos := start; pos <= len(x)-len(a); pos++ {
		match := true
		d := x[pos] - a[0]
		for offset := 1; offset < len(a); offset++ {
			if x[pos+offset]-a[offset] != d {
				match = false
				break
			}
		}
		if match {
			return pos
		}
	}
	return -1
}

func solve(x, a []int) []int {
	var (
		result []int
		start  int
	)
	for {
		pos := find(x, a, start)
		if pos == -1 {
			return result
		}
		result = append(result, pos+1)
		start = pos + 1
	}
}

func main() {
	scanner := makeScanner()
	_ = readInt(scanner)
	x := readArray(scanner)
	_ = readInt(scanner)
	a := readArray(scanner)
	printArray(solve(x, a))
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
	writer := bufio.NewWriter(os.Stdout)
	for i := 0; i < len(arr); i++ {
		writer.WriteString(strconv.Itoa(arr[i]))
		writer.WriteString(" ")
	}
	writer.Flush()
}
