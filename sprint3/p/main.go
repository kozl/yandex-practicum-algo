package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func solution(array []int) int {
	var min int
	part := []int{}
	result := 0
	for idx := range array {
		part = append(part, array[idx])
		curMin, curMax, continious := isContinious(part)
		if continious && curMin == min {
			result++
			part = []int{}
			min = curMax + 1
		}
	}
	return result
}

func isContinious(part []int) (min, max int, continious bool) {
	sort.Ints(part)
	min = part[0]
	max = part[len(part)-1]
	if len(part) == 1 {
		continious = true
		return
	}
	for idx := 0; idx < len(part)-1; idx++ {
		if part[idx+1]-part[idx] != 1 {
			continious = false
			return
		}
	}
	continious = true
	return
}

func main() {
	scanner := makeScanner()
	_ = readInt(scanner)
	array := readArray(scanner)
	result := solution(array)
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
