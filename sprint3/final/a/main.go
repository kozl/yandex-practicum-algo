package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solution(array []int, n int) int {
	left, right := 0, len(array)-1
	first := binarySearch(array, left, right, func(array []int, i int) int {
		if i-1 >= 0 && array[i] < array[i-1] {
			return 0
		}
		if array[i] >= array[0] {
			return 1
		}
		return -1
	})

	ordered := make([]int, len(array))
	if first == -1 {
		ordered = array
	} else {
		copy(ordered[:len(array)-first], array[first:])
		copy(ordered[len(array)-first:], array[:first])
	}

	pos := binarySearch(ordered, 0, len(ordered), func(array []int, i int) int {
		if array[i] == n {
			return 0
		}
		if array[i] < n {
			return 1
		}
		return -1
	})
	if first != -1 && pos < len(array)-first {
		return first + pos
	}
	return pos
}

func binarySearch(array []int, left, right int, fn func(array []int, i int) int) int {
	if left > right {
		return -1
	}
	mid := (left + right) / 2
	switch fn(array, mid) {
	case 0:
		return mid
	case 1:
		return binarySearch(array, mid+1, right, fn)
	case -1:
		return binarySearch(array, left, mid, fn)
	}
	return 0
}

func main() {
	scanner := makeScanner()
	_ = readInt(scanner)
	n := readInt(scanner)
	array := readArray(scanner)
	result := solution(array, n)
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
