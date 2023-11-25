package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func getTopN(array []int, n int) []int {
	m := map[int]int{}
	for _, el := range array {
		m[el]++
	}
	ids := [][2]int{}
	for k, v := range m {
		ids = append(ids, [2]int{k, v})
	}
	sort.Slice(ids, func(i, j int) bool {
		return ids[i][1] > ids[j][1]
	})

	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = ids[i][0]
	}
	return result
}

func main() {
	scanner := makeScanner()
	_ = readInt(scanner)
	array := readArray(scanner)
	n := readInt(scanner)
	result := getTopN(array, n)
	printArray(result)
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

func printArray(array []int) {
	if len(array) == 0 {
		return
	}
	s := make([]string, len(array))
	for idx, i := range array {
		s[idx] = strconv.Itoa(i)
	}
	fmt.Println(strings.Join(s, " "))
}
