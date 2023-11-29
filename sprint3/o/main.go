package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func getN(array []int, n int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(array)))
	left, right := minDiff(array), array[0]-array[len(array)-1]
	for left < right {
		mid := (left + right) / 2
		if checkIndex(array, mid, n) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func checkIndex(array []int, mid, n int) bool {
	var count int
	for i := 0; i < len(array)-1; i++ {
		var j int
		for j = i + 1; j < len(array); j++ {
			if array[i]-array[j] > mid {
				break
			}
		}
		count += j - i -1
	}
	return count >= n
}

func minDiff(array []int) int {
	min := array[0] - array[1]
	for i := 1; i < len(array)-1; i++ {
		if min > array[i]-array[i+1] {
			min = array[i] - array[i+1]
		}
	}

	return min
}

func main() {
	scanner := makeScanner()
	_ = readInt(scanner)
	array := readArray(scanner)
	n := readInt(scanner)
	result := getN(array, n)
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
