package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func modifyArrayNearestZeros(array []int) {
	currentZeroPos := -1
	var pos int

	for {
		for pos = currentZeroPos + 1; pos < len(array) && array[pos] != 0; pos++ {
			array[pos] = pos - currentZeroPos
		}
		if pos == len(array) {
			return
		}

		nextZeroPos := pos
		downTo := nextZeroPos - (nextZeroPos-currentZeroPos)/2
		if currentZeroPos == -1 {
			downTo = 0
		}
		for pos = nextZeroPos - 1; pos >= downTo; pos-- {
			array[pos] = nextZeroPos - pos
		}
		currentZeroPos = nextZeroPos
	}

}

func main() {
	scanner := makeScanner()
	_ = readInt(scanner)
	array := readArray(scanner)
	modifyArrayNearestZeros(array)
	printArray(array)
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 50 * 1024 * 1024
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
		_, _ = writer.WriteString(strconv.Itoa(arr[i]))
		_, _ = writer.WriteString(" ")
	}
	writer.Flush()
}
