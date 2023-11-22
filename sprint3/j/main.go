package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func bubbleSortAndPrint(w io.Writer, array []int) {
	var isPrinted bool
	for i := 0; i < len(array)-1; i++ {
		var changed bool
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] > array[j+1] {
				changed = true
				array[j+1], array[j] = array[j], array[j+1]
			}
		}
		if changed {
			isPrinted = true
			printArray(w, array)
		}
	}
	if !isPrinted {
		printArray(w, array)
	}
}

func printArray(w io.Writer, array []int) {
	s := make([]string, len(array))
	for i := 0; i < len(array); i++ {
		s[i] = strconv.Itoa(array[i])
	}
	fmt.Fprintln(w, strings.Join(s, " "))
}

func main() {
	scanner := makeScanner()
	_ = readInt(scanner)
	array := readArray(scanner)
	bubbleSortAndPrint(os.Stdout, array)
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 20 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func readInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	stringInt := scanner.Text()
	res, _ := strconv.Atoi(stringInt)
	return res
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
