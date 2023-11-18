package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func getNeighbours(matrix [][]int, row int, col int) []int {
	if row < 0 || row >= len(matrix) {
		return []int{}
	}
	if col < 0 || col >= len(matrix[row]) {
		return []int{}
	}
	neighboursInRow := []int{col - 1, col + 1}
	neighboursInCol := []int{row - 1, row + 1}
	result := []int{}
	for _, n := range neighboursInCol {
		if n < 0 || n >= len(matrix) {
			continue
		}
		result = append(result, matrix[n][col])
	}
	for _, n := range neighboursInRow {
		if n < 0 || n >= len(matrix[row]) {
			continue
		}
		result = append(result, matrix[row][n])
	}
	sort.Ints(result)
	return result
}

func main() {
	scanner := makeScanner()
	rows := readInt(scanner)
	cols := readInt(scanner)
	matrix := readMatrix(scanner, rows, cols)
	rowId := readInt(scanner)
	colId := readInt(scanner)
	neighbours := getNeighbours(matrix, rowId, colId)
	for _, elem := range neighbours {
		fmt.Print(elem)
		fmt.Print(" ")
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

func printArray(arr []int) {
	writer := bufio.NewWriter(os.Stdout)
	for i := 0; i < len(arr); i++ {
		writer.WriteString(strconv.Itoa(arr[i]))
		writer.WriteString(" ")
	}
	writer.Flush()
}

func readMatrix(scanner *bufio.Scanner, rows int, cols int) [][]int {
	matrix := make([][]int, rows)
	for i := 0; i < rows; i++ {
		matrix[i] = readArray(scanner)
	}
	return matrix
}
