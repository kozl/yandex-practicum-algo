package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func getN(array []int, n int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(array)))
	var k int
	n, k = calculateNK(len(array), n)

	diffs := []int{}

	for idx, v := range array {
		if idx+k < len(array) {
			diffs = append(diffs, v-array[idx+k])
		}
	}
	sort.Ints(diffs)
	return diffs[n-1]
}

func calculateNK(l, n int) (newN int, k int) {
	for k = 1; k <= l-1; k++ {
		if n-(l-k) <= 0 {
			break
		}
		n = n - (l - k)
	}
	return n, k
}

func addAndSort(diffs *list.List, value int, maxsize int) *list.List {
	if diffs.Back() != nil && diffs.Back().Value.(int) <= value && diffs.Len() == maxsize {
		return diffs
	}
	n := diffs.Front()
	for n != nil && n.Value.(int) <= value {
		n = n.Next()
	}
	if n == nil {
		if diffs.Len() < maxsize {
			diffs.PushBack(value)
		}
		return diffs
	}
	diffs.InsertBefore(value, n)
	if diffs.Len() > maxsize {
		diffs.Remove(diffs.Back())
	}
	return diffs
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
