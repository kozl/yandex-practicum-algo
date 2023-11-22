package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func customSort(array []int) {
	sort.Slice(array, func(i, j int) bool {
		a := strconv.Itoa(array[i])
		b := strconv.Itoa(array[j])

		x, _ := strconv.Atoi(a + b)
		y, _ := strconv.Atoi(b + a)

		return x > y
	})
}

func printArray(w io.Writer, array []int) {
	s := make([]string, len(array))
	for i := 0; i < len(array); i++ {
		s[i] = strconv.Itoa(array[i])
	}
	fmt.Fprintln(w, strings.Join(s, ""))
}

func main() {
	scanner := makeScanner()
	_ = readInt(scanner)
	array := readArray(scanner)
	customSort(array)
	printArray(os.Stdout, array)
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
