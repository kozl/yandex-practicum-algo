package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func transpose(matrix [][]int) [][]int {
	if len(matrix) == 0 {
		return [][]int{}
	}
	rows := len(matrix)
	cols := len(matrix[0])
	result := make([][]int, cols)
	for x, row := range matrix {
		for y, cell := range row {
			if result[y] == nil {
				result[y] = make([]int, rows)
			}
			result[y][x] = cell
		}
	}
	return result
}

func main() {
	scanner := makeScanner()
	rows := readInt(scanner)
	_ = readInt(scanner)
	matrix := make([][]int, rows)
	for i := 0; i < rows; i++ {
		arr := readArray(scanner)
		matrix[i] = arr
	}
	result := transpose(matrix)
	for _, row := range result {
		fmt.Println(strings.Join(toStringSlice(row), " "))
	}
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

func toStringSlice(s []int) []string {
	res := make([]string, len(s))
	for idx, i := range s {
		res[idx] = strconv.Itoa(i)
	}
	return res
}
